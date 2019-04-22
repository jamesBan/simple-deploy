package main

import (
	"net"
	"os"
	"log"
	"os/exec"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"errors"
	"strings"
	"golang.org/x/net/context"
	"time"
)

type project struct {
	Name       string `yaml:"name"`
	WorkerDir  string `yaml:"worker_dir"`
	ExeCommand string `yaml:"exec_command"`
	Timeout    int    `yaml:"timeout"`
}

type config struct {
	Server struct {
		Host     string   `yaml:"host"`
		Port     string   `yaml:"port"`
		AllowIps []string `yaml:"allow_ips"`
	}
	Projects []project `yaml:"project"`
}

var Config = config{}

func init() {
	configFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	defer configFile.Close()

	configContent, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err.Error())
	}

	err = yaml.Unmarshal(configContent, &Config)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	l, err := net.Listen("tcp", Config.Server.Host+":"+Config.Server.Port)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	log.Println("server listening:" + Config.Server.Host + ":" + Config.Server.Port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: " + err.Error())
			continue
		}

		clientIp := strings.Split(conn.RemoteAddr().String(), ":")[0]
		if !stringInSlice(clientIp, Config.Server.AllowIps) {
			log.Println("Error: not allow clien from" + clientIp)
			continue
		}

		go handleRequest(conn)
	}
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	readLen, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:" + err.Error())
		return
	}

	projectName := string(buf[:readLen])
	projectObj, err := checkProject(projectName)

	if err != nil {
		log.Println("Error:" + err.Error())
		return
	}
	log.Println("===========begin release: " + projectName + "=============")

	err = os.Chdir(projectObj.WorkerDir)
	if err != nil {
		log.Println(err.Error())
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(projectObj.Timeout)*time.Second)
	cmd := exec.CommandContext(ctx, projectObj.ExeCommand)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("command error: " + err.Error())
		return
	}
	log.Println("command out: " + string(out))

	log.Println("===========end release: " + projectName + "=============")
}

func checkProject(projectName string) (project, error) {
	for _, project := range Config.Projects {
		if project.Name == projectName {
			return project, nil
		}
	}

	return project{}, errors.New("invalid project:" + projectName)
}
