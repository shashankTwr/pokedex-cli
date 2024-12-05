package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
const baseURL = "https://pokeapi.co/api/v2/location-area/"

var cfg = &config{
    Count: 0,
    Next: "",
    Previous: "",
    Results: nil,
}

type config struct {
    Count    int    `json:"count"`
    Next     string `json:"next"`
    Previous string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

// GetNextLocations handles the "map" command
func GetNextLocations() error {
    currentURL := baseURL
    if cfg.Next != "" {
        currentURL = cfg.Next
    }

    return getAndPrintLocations(currentURL)
}

// GetPreviousLocations handles the "mapb" command
func GetPreviousLocations() error {
    if cfg.Previous == "" {
        return fmt.Errorf("this is the first page")
    }
    
    return getAndPrintLocations(cfg.Previous)
}

// private helper function to avoid code duplication
func getAndPrintLocations(url string) error {
    res, err := http.Get(url)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    dat, err := io.ReadAll(res.Body)
    if err != nil {
        return err
    }

    c := config{}
    err = json.Unmarshal(dat, &c)
    if err != nil {
        return err
    }

    for _, result := range c.Results {
        fmt.Println(result.Name)
    }
    cfg = &c
    return nil
}