package main

import (
	"fmt"
	"strconv"
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
	firstURLWords := strings.Split(locations.Results[0].URL, "/")
	lastURLWords := strings.Split(locations.Results[len(locations.Results)-1].URL, "/")

	firstPosition, err := strconv.Atoi(firstURLWords[len(firstURLWords)-2])
	if err != nil {
		return fmt.Errorf("error converting location position to int: %w", err)
	}

	fmt.Printf("\nShowing from %d to %s of %d locations", firstPosition, lastURLWords[len(lastURLWords)-2], locations.Count)

	for idx, location := range locations.Results {

		position := idx + firstPosition
		fmt.Printf("\n %d - %s.", position, location.Name)
	}

	fmt.Printf("\n\n")
	return nil
}
