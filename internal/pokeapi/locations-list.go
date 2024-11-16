package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	locationsResponse := RespShallowLocations{}

	url := BASE_URL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return locationsResponse, fmt.Errorf("error creating the request: %w", err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&locationsResponse); err != nil {
		return locationsResponse, fmt.Errorf("error decoding response body: %w", err)
	}

	return locationsResponse, nil
}
