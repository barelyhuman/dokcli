package main

import (
	"io/ioutil"
	"log"
)

func main() {
	log.Println("Reading Config")

	config, err := readConfig()
	if err != nil {
		log.Fatal("Error reading config...\n", err)
	}

	log.Println("Generating Script")
	script := config.GenerateScript()
	log.Println("Generating Domain Script")
	domainScript := config.GenerateDomainScript()

	scriptFilePath := "./dokku-setup-" + config.App.Name + ".sh"
	domainScriptFilePath := "./dokku-setup-" + config.App.Name + "-domain.sh"

	err = ioutil.WriteFile(scriptFilePath, []byte(script), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(domainScriptFilePath, []byte(domainScript), 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Generated Scripts:%s, %s", scriptFilePath, domainScriptFilePath)

}
