package controller

import (
	"fmt"
	"github.com/Kudzeri/Boot.dev-pokedex-go/model"
	"os"
	"strings"
)

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range model.CliCommandList {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}

	return nil
}

func CommandExit() error {
	os.Exit(0)
	return nil
}

func ExecuteCommand(commandName string) {
	commandName = strings.TrimSpace(commandName)
	if cmd, exists := model.CliCommandList[commandName]; exists {
		err := cmd.Callback()
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	} else {
		fmt.Println("Command not found:", commandName)
		fmt.Println(model.CliCommandList)
	}
}

func init() {
	model.CliCommandList = map[string]model.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
	}
}
