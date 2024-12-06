package pokeapi

import (
	"net/http"
	"time"

	"github.com/shashankTwr/pokedexcli/internal/pokecache"
)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

// New Client

func NewClient(timeout, cacheInterval time.Duration) Client{
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}