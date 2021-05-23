package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v2"
)

// AppConfig the base app config for creating the dokku script
type AppConfig struct {
	Plugins map[string]string
	App     struct {
		Name             string
		DB               string
		DBName           string `yaml:"dbName"`
		Domain           string
		HTTPS            bool
		LetsEncryptEmail string `yaml:"letsEncryptEmail"`
	}
}

const (
	dokku             = "dokku"
	createCmd         = ":create"
	linkCmd           = ":link"
	pluginInstall     = "sudo " + dokku + " plugin:install"
	domainAdd         = dokku + " domains:add"
	configAdd         = dokku + " configs:set"
	letsencryptEnable = dokku + " letsencrypt:enable"
)

var pluginURLs = map[string]string{
	"mongo":       "https://github.com/dokku/dokku-mongo.git",
	"postgres":    "https://github.com/dokku/dokku-postgres.git",
	"letsencrypt": "https://github.com/dokku/dokku-letsencrypt.git",
}

var configQuestions = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "Name of the app?"},
		Validate: survey.Required,
	},
	{
		Name:     "dbName",
		Prompt:   &survey.Input{Message: "Name of the database?"},
		Validate: survey.Required,
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
	{
		Name:   "https",
		Prompt: &survey.Confirm{Message: "You want to enable HTTPS (add LetsEncrypt)?"},
	},
}

var letsEncryptEmailQuestion = &survey.Input{
	Message: "What email do you want to use with LetsEncrypt?",
}

// {
// 	Name:   "letsEncryptEmail",
// 	Prompt: &survey.Input{Message: "What email do you want to use with LetsEncrypt?"},
// }

func (config *AppConfig) createApp() string {
	return dokku + " apps" + createCmd + " " + config.App.Name + "\n"
}

func (config *AppConfig) addDomain() string {
	domain := stripProtocol(config.App.Domain)
	return domainAdd + " " + config.App.Name + " " + domain + "\n"
}

func (config *AppConfig) addLetsEncrypt() string {

	var sb strings.Builder

	if config.App.HTTPS && config.App.LetsEncryptEmail == "" {
		log.Fatal("`letsEncryptEmail:` is needed if you set http as true")
	}

	sb.Write([]byte(configAdd + " --no-restart " + config.App.Name + " DOKKU_LETSENCRYPT_EMAIL=" + config.App.LetsEncryptEmail + "\n"))
	sb.Write([]byte(letsencryptEnable + " " + config.App.Name))

	return sb.String()
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
	return sb.String()
}

func (config *AppConfig) GenerateDomainScript() string {
	var sb strings.Builder
	sb.Write([]byte(config.addDomain()))
	if config.App.HTTPS {
		sb.Write([]byte(config.addLetsEncrypt()))
	}
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
		Name             string
		DB               string `survey:"db"`
		DBName           string `survey:"dbName"`
		Domain           string
		HTTPS            bool
		LetsEncryptEmail string `survey:"letsEncryptEmail"`
	}{}

	err := survey.Ask(configQuestions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	config.App.Name = answers.Name
	config.App.DBName = answers.DBName
	config.App.Domain = answers.Domain
	config.App.HTTPS = answers.HTTPS

	if answers.HTTPS {
		err = survey.AskOne(letsEncryptEmailQuestion, &answers.LetsEncryptEmail)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		config.App.LetsEncryptEmail = answers.LetsEncryptEmail
	}

	pluginNameLower := strings.ToLower(answers.DB)

	config.App.DB = pluginNameLower
	pluginsMap := make(map[string]string)

	pluginsMap[pluginNameLower] = pluginURLs[pluginNameLower]

	if answers.HTTPS {
		pluginsMap["letsencrypt"] = pluginURLs["letsencrypt"]
	}

	config.Plugins = pluginsMap

	return nil
}

func stripProtocol(domain string) string {
	parsedURL, err := url.Parse(domain)
	if err != nil {
		log.Fatal(err)
	}

	if parsedURL.Scheme == "" {
		return stripProtocol("https://" + domain)
	}

	return parsedURL.Host
}
