package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := BASE_URL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if value, ok := c.cache.Get(url); ok {
		locationsResponse := RespShallowLocations{}
		err := json.Unmarshal(value, &locationsResponse)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResponse, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("error creating the request: %w", err)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("error reading response: %w", err)
	}

	locationsResponse := RespShallowLocations{}
	if err := json.Unmarshal(data, &locationsResponse); err != nil {
		return RespShallowLocations{}, fmt.Errorf("error decoding response body: %w", err)
	}

	c.cache.Add(url, data)
	return locationsResponse, nil
}
