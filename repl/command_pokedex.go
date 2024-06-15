package repl

import (
	"errors"
	"fmt"
	"github.com/Kudzeri/Boot.dev-pokedex-go/repl/config"
)

func CallbackPokedex(cfg *config.Config, args ...string) error {
	if len(cfg.CaughtPokemon) == 0 {
		return errors.New("you have 0 pokemon in pokedex")
	}

	fmt.Println("Pokemon in Pokedex:")
	for _, pokemon := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
