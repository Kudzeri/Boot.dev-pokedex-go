package repl

import (
	"fmt"
	"github.com/Kudzeri/Boot.dev-pokedex-go/repl/config"
)

func CallbackHelp(cfg *config.Config) error {
	fmt.Println("Welcome to pokedex!")
	fmt.Println("This list of command:")
	commands := GetCommands()
	for _, cmd := range commands {
		fmt.Printf("-%v: %v\n", cmd.Name, cmd.Description)
	}

	fmt.Println()
	return nil
}
