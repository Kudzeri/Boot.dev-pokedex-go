package repl

import (
	"errors"
	"fmt"
	"github.com/Kudzeri/Boot.dev-pokedex-go/repl/config"
	"math/rand"
)

func CallbackCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	fmt.Printf("%s was caught!", pokemonName)
	return nil
}
