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

		fmt.Println(text)
	}

}

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
