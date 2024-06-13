package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex> ")

		scanner.Scan()
		text := CleanInput(scanner.Text())

		if len(text) == 0 {
			continue
		}
		commandName := text[0]
		availableCommands := GetCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		command.Callback()
	}
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func()
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Prints the help menu",
			Callback:    CallbackHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Turns off the Pokedex",
			Callback:    CallbackExit,
		},
	}
}

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
