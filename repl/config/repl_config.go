package config

import "github.com/Kudzeri/Boot.dev-pokedex-go/internal/pokeapi"

type Config struct {
	PokeapiClient       pokeapi.Client
	NextLocationAreaURL *string
	PrevLocationAreaURL *string
}
