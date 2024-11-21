package main

import (
	"fmt"
	"strings"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

func cmdMap(cfg *config.TConfig, args ...string) error {
	locationsResponse, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	err = logLocations(locationsResponse)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResponse.Next
	cfg.PrevLocationsURL = locationsResponse.Previous

	return nil
}

func cmdMapB(cfg *config.TConfig, args ...string) error {
	if cfg.PrevLocationsURL == nil {
		return fmt.Errorf("you're on the first page cannot navigate back")
	}
	locationsResponse, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	err = logLocations(locationsResponse)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResponse.Next
	cfg.PrevLocationsURL = locationsResponse.Previous

	return nil
}

func logLocations(locations pokeapi.RespShallowLocations) error {
	firstPosition := getPositionFromURL(locations.Results[0].URL)
	lastPosition := getPositionFromURL(locations.Results[len(locations.Results)-1].URL)

	fmt.Printf("\nShowing from %s to %s of %d locations", firstPosition, lastPosition, locations.Count)

	for _, location := range locations.Results {
		position := getPositionFromURL(location.URL)
		fmt.Printf("\n %04s - %s.", position, location.Name)
	}

	fmt.Printf("\n\n")
	return nil
}

func getPositionFromURL(url string) string {
	words := strings.Split(url, "/")
	return words[len(words)-2]
}
