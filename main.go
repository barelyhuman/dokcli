// SPDX-License-Identifier: MIT

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type TNewApp struct {
	name     string
	database string
}

type Command struct {
}

func (command *Command) Run() error {
	return nil
}

func main() {
	var newApp = flag.NewFlagSet("new", flag.ExitOnError)
	var appName = newApp.String("name", "", "Name of the new app")

	if len(os.Args) < 2 {
		fmt.Println("expected one of the subcommands")
		os.Exit(1)
	}

	command := deletgateToCommand()

	err := command.Run()

	if err.err != nil {
		log.Fatal(err.message, err.err)
	}
}

func deletgateToCommand() Command {
	// Decide which command to deletgate the operation to
}
