package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/shashankTwr/pokedexcli/internal/pokecache"
)

// GetNextLocations handles the "map" command
func GetNextLocations(cache *pokecache.Cache, cfg *Config, baseURL string) error {
    currentURL := baseURL
    if cfg.Next != "" {
        currentURL = cfg.Next
    }

    return getAndPrintLocations(cache, cfg, currentURL)
}

// GetPreviousLocations handles the "mapb" command
func GetPreviousLocations(cache *pokecache.Cache, cfg *Config) error {
    if cfg.Previous == "" {
        return fmt.Errorf("this is the first page")
    }
    
    return getAndPrintLocations(cache, cfg ,cfg.Previous)
}

// private helper function to avoid code duplication
func getAndPrintLocations(cache *pokecache.Cache,cfg *Config,url string) error {


    dat, err := fetchPokemonData(cache, url)
    if err!= nil{
        return err
    }

    err = json.Unmarshal(dat, &cfg)
    if err != nil {
        return err
    }

    for _, result := range cfg.Results {
        fmt.Println(result.Name)
    }
    return nil
}

func fetchPokemonData(cache *pokecache.Cache, url string) ([]byte, error) {
	if data, found := cache.Get(url); found {
		// fmt.Println("Cache hit")
		return data, nil
	}
	// if not in cache make the http request
	// fmt.Println("Cache miss, making request")
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	cache.Add(url, data)
	return data, nil
}
