package main

import (
	"github.com/Kudzeri/Boot.dev-pokedex-go/internal/pokeapi"
	"github.com/Kudzeri/Boot.dev-pokedex-go/repl"
	"github.com/Kudzeri/Boot.dev-pokedex-go/repl/config"
	"time"
)

func main() {

	cfg := config.Config{
		PokeapiClient: pokeapi.NewClient(time.Hour),
	}
	repl.StartRepl(&cfg)
}
