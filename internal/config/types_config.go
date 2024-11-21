package config

import "github.com/ingcapadev/pokedex-with-go/internal/pokeapi"

type TConfig struct {
	PokeapiClient    pokeapi.Client `json:"-"`
	CaughtPokemon    map[string]pokeapi.Pokemon
	NextLocationsURL *string
	PrevLocationsURL *string
}
