package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type HookConfig struct {
	Path   string    `json:"path"`
	Secret string    `json:"secret"`
	Log    LogConfig `json:"log"`
}

type LogConfig struct {
	File  string `json:"file"`
	Level string `json:"level"`
	Mode  string `json:"mode"`
}

const (
	defaultContextPath = "/webhooks/gitlab/"

	defaultLogFile  = "./rabbit-hook.log"
	defaultLogLevel = "INFO"
	defaultLogMode  = "rabbit"

	pathSuffix = "/"
)

func readConfig(configPath string) (HookConfig, error) {
	fileData, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		panic(err)
	}

	config := HookConfig{}
	json.Unmarshal(fileData, &config)
	return config, err
}

func checkConfig(config HookConfig) {
	if config.Path == "" {
		config.Path = defaultContextPath
	} else {
		if !strings.HasSuffix(config.Path, pathSuffix) {
			config.Path = config.Path + pathSuffix
		}
	}
	if config.Log.File == "" {
		config.Log.File = defaultLogFile
	}

	if config.Log.Level == "" {
		config.Log.Level = defaultLogLevel
	}

	if config.Log.Mode == "" {
		config.Log.Mode = defaultLogMode
	}
}
