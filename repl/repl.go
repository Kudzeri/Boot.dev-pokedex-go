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
		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}
		availableCommands := GetCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.Callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config, ...string) error
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
			Description: "Prints pokedex location areas",
			Callback:    CallbackMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Back the previous page of location areas",
			Callback:    CallbackMapb,
		},
		"explore": {
			Name:        "explore {location_area}",
			Description: "List the pokemon in a location area",
			Callback:    CallbackExplore,
		},
		"catch": {
			Name:        "catch {pokemon_name}",
			Description: "Try to catch the pokemon",
			Callback:    CallbackCatch,
		},
	}
}

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
