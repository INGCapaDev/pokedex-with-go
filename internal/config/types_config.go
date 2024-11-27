package config

import "github.com/ingcapadev/pokedex-with-go/internal/pokeapi"

type TConfig struct {
	PokeapiClient    pokeapi.Client
	CaughtPokemon    map[string]pokeapi.Pokemon
	Inventory        Inventory
	NextLocationsURL *string
	PrevLocationsURL *string
}
