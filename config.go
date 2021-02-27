package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v2"
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

var supportedDatabases = map[string]string{
	"mongo":    "https://github.com/dokku/dokku-mongo.git",
	"postgres": "https://github.com/dokku/dokku-postgres.git",
}

var configQuestions = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "Name of the app?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:      "dbName",
		Prompt:    &survey.Input{Message: "Name of the database?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "db",
		Prompt: &survey.Select{
			Message: "Select a db plugin:",
			Options: []string{"Mongo", "Postgres"},
			Default: "Postgres",
		},
	},
	{
		Name:   "domain",
		Prompt: &survey.Input{Message: "What domain do you want this app to have?"},
	},
}

var addMorePrompt = &survey.Confirm{
	Message: "Add more ?",
}

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
	fromFile := true
	log.Println("Looking for `dokku-gen.yml`")
	data, err := ioutil.ReadFile("./dokku-gen.yml")
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			fromFile = false
			askConfigQuestions(&config)
		}
	}

	if fromFile {
		log.Println("Decoding Config")
		err = yaml.Unmarshal([]byte(data), &config)
		if err != nil {
			return config, err
		}
	}

	return config, nil
}

func askConfigQuestions(config *AppConfig) error {
	answers := struct {
		Name   string
		DB     string `survey:"db"`
		DBName string `survey:"dbName"`
		Domain string
	}{}

	err := survey.Ask(configQuestions, &answers)

	config.App.Name = answers.Name
	config.App.DBName = answers.DBName
	config.App.Domain = answers.Domain

	pluginNameLower := strings.ToLower(answers.DB)

	config.App.DB = pluginNameLower
	pluginsMap := make(map[string]string)

	pluginsMap[pluginNameLower] = supportedDatabases[pluginNameLower]

	config.Plugins = pluginsMap

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
