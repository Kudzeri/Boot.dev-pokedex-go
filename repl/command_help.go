package repl

import (
	"fmt"
)

func CallbackHelp() {
	fmt.Println("Welcome to pokedex!")
	fmt.Println("This list of command:")
	commands := GetCommands()
	for _, cmd := range commands {
		fmt.Printf("-%v: %v\n", cmd.Name, cmd.Description)
	}

}