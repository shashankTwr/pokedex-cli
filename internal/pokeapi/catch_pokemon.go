package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName *string) (RespPokemon, error) {
    url := baseURL + "/pokemon/" + *pokemonName

	if val, ok := c.cache.Get(url); ok {
        pokemonResp := RespPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespPokemon{}, errors.New("failed to unmarshal pokemon in CatchPokemon")
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, errors.New("failed to create request in CatchPokemon")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, errors.New("failed to send response in CatchPokemon")
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, errors.New("failed to read resp.body in CatchPokemon")
	}

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, errors.New("failed to read dat in CatchPokemon")
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}