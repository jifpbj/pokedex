package pokeapi

import (
	"net/http"
	"time"

	"github.com/jifpbj/pokedex/internal/pokecache"
)

// Client ....
type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

// NewClient creates a new PokeAPI client with a custom timeout.
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
