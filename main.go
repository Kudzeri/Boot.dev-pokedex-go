package main

import (
	"bufio"
	"fmt"
	"github.com/Kudzeri/Boot.dev-pokedex-go/controller"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("pokedex > ")

		input, err := sc.ReadString('\n')
		if err != nil {
			fmt.Println("Reading error:", err)
			return
		}

		controller.ExecuteCommand(input)
	}
}
