package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocations(location_area *string) (RespExploreLocations, error) {
    url := baseURL + "/location-area/" + *location_area

	if val, ok := c.cache.Get(url); ok {
        locationsResp := RespExploreLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespExploreLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploreLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploreLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreLocations{}, err
	}

	locationsResp := RespExploreLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespExploreLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}