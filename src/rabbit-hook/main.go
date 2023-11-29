package main

import (
	"flag"
	"fmt"
	"github.com/op/go-logging"
	"gopkg.in/go-playground/webhooks.v5/gitlab"
	"net/http"
	"os"
	"strings"
)

const (
	defaultConfigFile = "config/application.json"
	defaultPort       = "8077"
	defaultHost       = ":"
)

func main() {
	var configPath string
	var port string = ":"
	flag.StringVar(&configPath, "config", defaultConfigFile, "define a config path, default is config/application.json")
	flag.StringVar(&port, "port", defaultPort, "define listen port, default is 8077")
	flag.Parse()
	host := defaultHost + port
	config, configErr := readConfig(configPath)
	if configErr != nil {
		os.Exit(1)
	}
	checkConfig(config)

	Init(config.Log.Level, config.Log.File, config.Log.Mode)
	log := logging.MustGetLogger(config.Log.Mode)

	hook, _ := gitlab.New(gitlab.Options.Secret(config.Secret))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "a huge rabbit!")
	})

	http.HandleFunc(config.Path, func(w http.ResponseWriter, r *http.Request) {
		robotId := strings.TrimPrefix(r.URL.Path, config.Path)
		if robotId == "" {
			log.Warning("undefined robot, stop process")
			return
		}
		payload, err := hook.Parse(r, gitlab.PushEvents, gitlab.MergeRequestEvents, gitlab.IssuesEvents)
		if err != nil {
			if err == gitlab.ErrEventNotFound {
				log.Warning(err)
			} else {
				log.Error(err)
			}
		}
		switch payload.(type) {

		case gitlab.PushEventPayload:
			release := payload.(gitlab.PushEventPayload)
			pushMessage := buildPushNotify(release)
			sendMessage(robotId, pushMessage, config.Log.Mode)

		case gitlab.IssueEventPayload:
			issue := payload.(gitlab.IssueEventPayload)
			issueMessage := buildIssueNotify(issue)
			sendMessage(robotId, issueMessage, config.Log.Mode)

		case gitlab.MergeRequestEventPayload:
			mergeRequest := payload.(gitlab.MergeRequestEventPayload)
			mergeRequestMessage := buildMergeRequestNotify(mergeRequest)
			sendMessage(robotId, mergeRequestMessage, config.Log.Mode)
		}
	})

	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
