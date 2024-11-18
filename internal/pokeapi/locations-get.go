package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (Location, error) {
	url := BASE_URL + "/location-area/" + name

	if value, ok := c.cache.Get(url); ok {
		locationResponse := Location{}
		err := json.Unmarshal(value, &locationResponse)
		if err != nil {
			return Location{}, err
		}
		return locationResponse, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Location{}, fmt.Errorf("error making the request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return Location{}, fmt.Errorf("invalid location: location %s not found", name)
		}
		return Location{}, fmt.Errorf("error non-OK HTTP status code: %d", res.StatusCode)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("error reading response: %w", err)
	}

	locationResponse := Location{}
	if err := json.Unmarshal(data, &locationResponse); err != nil {
		return Location{}, fmt.Errorf("error decoding response body: %w", err)
	}

	c.cache.Add(url, data)
	return locationResponse, nil
}
