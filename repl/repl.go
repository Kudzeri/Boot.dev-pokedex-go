package repl

import (
	"bufio"
	"fmt"
	"github.com/Kudzeri/Boot.dev-pokedex-go/repl/config"
	"os"
	"strings"
)

func StartRepl(cfg *config.Config) {
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

		command.Callback(cfg)
	}
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config) error
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
		"map": {
			Name:        "map",
			Description: "Prints pokedex map locations",
			Callback:    CallbackMap,
		},
	}
}

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
