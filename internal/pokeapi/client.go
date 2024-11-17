package pokeapi

import (
	"net/http"
	"time"

	"github.com/ingcapadev/pokedex-with-go/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration, cacheLifetime time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheLifetime),
	}
}
