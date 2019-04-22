package main

import (
	"net/http"
	"io/ioutil"
	"os"
	"sync"
	"gopkg.in/yaml.v2"
)

var wg sync.WaitGroup
var Config = ServerConfig{}

//parse config
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello"))
	})

	mux.Handle("/release/project", &HandleProject{config:Config})

	logMessage("server start listen " + Config.Server.Host + ":" + Config.Server.Port)
	err := http.ListenAndServe(Config.Server.Host+":"+Config.Server.Port, mux)
	if err != nil {
		panic(err.Error())
	}
}
