package main

import (
	"encoding/json"
	"fmt"
	"github.com/op/go-logging"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	messageType       = "markdown"
	assignMessageType = "text"

	assignMessageTemplate = "#%d assign to: "

	sendUrlPattern = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
)

type markdownRequest struct {
	Msgtype  string   `json:"msgtype"`
	Markdown markdown `json:"markdown"`
}

type textRequest struct {
	Msgtype string `json:"msgtype"`
	Text    text   `json:"text"`
}

type text struct {
	Content       string   `json:"content"`
	MentionedList []string `json:"mentioned_list"`
}

type markdown struct {
	Content string `json:"content"`
}

func sendMessage(robotId string, message string, logMode string) {
	log := logging.MustGetLogger(logMode)
	url := fmt.Sprintf(sendUrlPattern, robotId)

	var content markdown = markdown{message}
	var requestBodyModel markdownRequest = markdownRequest{messageType, content}

	requestBody, err := json.Marshal(requestBodyModel)
	log.Info(string(requestBody))
	if err != nil {
		log.Error("parse request to json string failed, stop process!")
		return
	}

	response, err := http.Post(url, "application/json; charset=UTF-8", strings.NewReader(string(requestBody)))
	if err != nil {
		log.Warning(fmt.Sprintf("send message failed, reason: %s", err))
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		log.Info(fmt.Sprintf("request success! code: %d, body: %s", response.StatusCode, body))
	}
}

func sendAssign(robotId string, mentioned []string, iid int64, logMode string) {
	log := logging.MustGetLogger(logMode)
	url := fmt.Sprintf(sendUrlPattern, robotId)

	message := fmt.Sprintf(assignMessageTemplate, iid)
	var content text = text{message, mentioned}
	var requestBodyModel textRequest = textRequest{assignMessageType, content}

	requestBody, err := json.Marshal(requestBodyModel)
	log.Info(string(requestBody))
	if err != nil {
		log.Error("parse request to json string failed, stop process!")
		return
	}

	response, err := http.Post(url, "application/json; charset=UTF-8", strings.NewReader(string(requestBody)))
	if err != nil {
		log.Warning(fmt.Sprintf("send message failed, reason: %s", err))
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		log.Info(fmt.Sprintf("request success! code: %d, body: %s", response.StatusCode, body))
	}
}
