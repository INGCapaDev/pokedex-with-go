package main

import (
	"testing"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

func TestInspectCMD(t *testing.T) {
	mockCaughtPokemon := map[string]pokeapi.Pokemon{
		"pikachu":   {},
		"bulbasaur": {},
		"squirtle":  {},
	}

	mockConfig := &config.TConfig{
		PokeapiClient: pokeapi.Client{},
		CaughtPokemon: mockCaughtPokemon,
	}

	err := cmdInspect(mockConfig)
	if err == nil {
		t.Errorf("expected to throw an error if no pokemon is passed")
	}

	err = cmdInspect(mockConfig, "bulbasaur", "squirtle")
	if err == nil {
		t.Errorf("expected to throw an error if more than one pokemon is passed")
	}

	err = cmdInspect(mockConfig, "pidgey")
	if err == nil {
		t.Errorf("expected to throw an error for uncaught pokemon")
	}

	err = cmdInspect(mockConfig, "pikachu")
	if err != nil {
		t.Errorf("expected to print pokemon stats")
	}
}
