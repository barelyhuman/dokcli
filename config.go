package main

import (
	"io/ioutil"
	"log"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// AppConfig the base app config for creating the dokku script
type AppConfig struct {
	Plugins map[string]string
	App     struct {
		Name   string
		DB     string
		DBName string `yaml:"dbName"`
		Domain string
	}
}

const (
	dokku         = "dokku"
	createCmd     = ":create"
	linkCmd       = ":link"
	pluginInstall = "sudo " + dokku + " plugin:install"
	domainAdd     = dokku + " domains:add"
)

func (config *AppConfig) createApp() string {
	return dokku + " apps" + createCmd + " " + config.App.Name + "\n"
}

func (config *AppConfig) addDomain() string {
	return domainAdd + " " + config.App.Name + " " + config.App.Domain + "\n"
}

func (config *AppConfig) installPlugins() string {
	var sb strings.Builder

	for k := range config.Plugins {
		sb.Write([]byte(pluginInstall + " " + config.Plugins[k] + "\n"))
	}

	return sb.String()
}

func (config *AppConfig) createDatabase() string {
	return dokku + " " + config.App.DB + createCmd + " " + config.App.DBName + "\n"
}

func (config *AppConfig) linkDatabase() string {
	return dokku + " " + config.App.DB + linkCmd + " " + config.App.DBName + " " + config.App.Name + "\n"
}

// GenerateScript - generate a bash script with dokku commands in flow
func (config *AppConfig) GenerateScript() string {
	var sb strings.Builder
	sb.Write([]byte("#!/bin/bash\n\n"))
	sb.Write([]byte(config.installPlugins()))
	sb.Write([]byte(config.createApp()))
	sb.Write([]byte(config.createDatabase()))
	sb.Write([]byte(config.linkDatabase()))
	sb.Write([]byte(config.addDomain()))
	return sb.String()
}

func readConfig() (AppConfig, error) {
	config := AppConfig{}
	log.Println("Looking for `dokku-gen.yml`")
	data, err := ioutil.ReadFile("./dokku-gen.yml")
	if err != nil {
		return config, err
	}

	log.Println("Decoding Config")
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// Read and use config parameters from an ini config reader
