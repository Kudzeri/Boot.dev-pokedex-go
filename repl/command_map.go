package repl

import (
	"fmt"
	"github.com/Kudzeri/Boot.dev-pokedex-go/internal/pokeapi"
	"github.com/Kudzeri/Boot.dev-pokedex-go/repl/config"
	"log"
)

func CallbackMap(cfg *config.Config) error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}
