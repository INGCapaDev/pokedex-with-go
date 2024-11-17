package main

import (
	"fmt"
	"strings"

	"github.com/ingcapadev/pokedex-with-go/internal/pokeapi"
)

func cmdMap(cfg *config, args ...string) error {
	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	err = logLocations(locationsResponse)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResponse.Next
	cfg.prevLocationsURL = locationsResponse.Previous

	return nil
}

func cmdMapB(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page cannot navigate back")
	}
	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	err = logLocations(locationsResponse)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResponse.Next
	cfg.prevLocationsURL = locationsResponse.Previous

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
