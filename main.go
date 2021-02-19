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

	scriptFilePath := "./dokku-setup-" + config.App.Name + ".sh"

	err = ioutil.WriteFile(scriptFilePath, []byte(script), 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Generated:%s", scriptFilePath)

}
