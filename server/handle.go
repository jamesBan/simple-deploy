package main

import (
	"io/ioutil"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net"
	"os"
	"time"
	"net/http"
	"errors"
	"log"
	"fmt"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type HandleProject struct {
	config      ServerConfig
	requestBody []byte
	request     *http.Request
	response    http.ResponseWriter
}

func (this *HandleProject) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.request = r
	this.response = w

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		this.convertReponse(400, "empty body")
		return
	}

	this.requestBody = requestBody

	if code, err := this.checkSign(); err != nil {
		this.convertReponse(code, err.Error())
		return
	}

	body, err := this.parseJsonData()
	if err != nil {
		this.convertReponse(400, err.Error())
		return
	}

	//releaseTitle := body.Release.Name
	//releaseBody := body.Release.Body
	//releaseDraft := body.Release.Draft
	projectName := body.Repository.Name
	ProjectList := this.config.ReleaseServer

	if _, ok := ProjectList[projectName]; !ok {
		this.convertReponse(400, "error project")
		return
	}

	logMessage("==========begin release: " + projectName + "================")
	logMessage("title:" + body.Release.Name)
	for i := 0; i < len(ProjectList[projectName]); i++ {
		wg.Add(1)
		go this.sendCommand(ProjectList[projectName][i], projectName)
	}
	wg.Wait()

	logMessage("==========end release: " + projectName + "================")
	logMessage("release success!")
}

func (this *HandleProject) convertReponse(code int, message string) {
	this.response.WriteHeader(code)
	this.response.Write([]byte(message))
}

//检查头
func (this *HandleProject) checkHead() (int, error) {
	if this.request.Header.Get("X-Gogs-Event") != "release" || len(this.request.Header.Get("X-Gogs-Signature")) < 1 {
		return 400, errors.New("empty head event")
	}

	return 200, nil
}

//检查签名
func (this *HandleProject) checkSign() (int, error) {
	headSign := this.request.Header.Get("X-Gogs-Signature")
	generateHash := this.signHash(this.requestBody)

	if headSign != generateHash {
		return 400, errors.New("bad signature")
	}

	return 200, nil
}

//计算签名
func (this *HandleProject) signHash(message []byte) (string) {
	hash := hmac.New(sha256.New, []byte(this.config.Server.Secret))
	hash.Write(message)
	return hex.EncodeToString(hash.Sum(nil))
}

//解析json
func (this *HandleProject) parseJsonData() (JsonData, error) {
	body := JsonData{}
	err := json.Unmarshal(this.requestBody, &body)
	return body, err
}

func (this *HandleProject) sendCommand(serverIp, projectName string) {
	defer wg.Done()
	defer logMessage("end release: " + serverIp)

	logMessage("begin release: " + serverIp)
	client, err := net.Dial("tcp", serverIp)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer client.Close()

	_, err = client.Write([]byte(projectName))
	if err != nil {
		log.Println(err.Error())
	}
}

func logMessage(message string) {
	fmt.Fprintln(os.Stdout, fmt.Sprintf("%s：%s", time.Now().String(), message))
}
