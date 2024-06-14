package repl

import (
	"errors"
	"fmt"
	"github.com/Kudzeri/Boot.dev-pokedex-go/repl/config"
)

func CallbackMap(cfg *config.Config) error {

	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.NextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.NextLocationAreaURL = resp.Next
	cfg.PrevLocationAreaURL = resp.Previous

	return nil
}

func CallbackMapb(cfg *config.Config) error {
	if cfg.PrevLocationAreaURL == nil {
		return errors.New("you're on first page")
	}
	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.PrevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.NextLocationAreaURL = resp.Next
	cfg.PrevLocationAreaURL = resp.Previous

	return nil
}
