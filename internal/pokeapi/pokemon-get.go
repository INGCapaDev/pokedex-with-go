package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	if name == "" {
		return Pokemon{}, fmt.Errorf("you must provide a valid pokemon name")
	}

	url := BASE_URL + "/pokemon/" + name

	if value, ok := c.cache.Get(url); ok {
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(value, &pokemonResponse)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResponse, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error making the request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return Pokemon{}, fmt.Errorf("invalid pokemon name: %s not found", name)
		}
		return Pokemon{}, fmt.Errorf("error non-OK HTTP status code: %d", res.StatusCode)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response: %w", err)
	}

	pokemonResponse := Pokemon{}
	if err := json.Unmarshal(data, &pokemonResponse); err != nil {
		return Pokemon{}, fmt.Errorf("error decoding response body %w", err)
	}

	return pokemonResponse, nil
}
