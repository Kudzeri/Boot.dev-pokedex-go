package main

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
		command := text[0]

		switch command {
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Println("Welcome to pokedex!")
			fmt.Println("This list of command:")
			fmt.Println("- help")
			fmt.Println("- exit")
		default:
			fmt.Println("invalid command")

		}
	}

}

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
