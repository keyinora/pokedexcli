package pokeapi

import (
	"encoding/json"
	"errors"
)

func (c *Client) ExploreLocation(area_name string) (RespLocationEncounters, error) {
	if area_name == "" {
		return RespLocationEncounters{}, errors.New("No area name given")
	}

	url := baseURL + "/location-area/" + area_name

	// Step 1: check cache
	if data, found := cache.Get(url); found {
		// If cached data is found, unmarshal it and return
		resp := RespLocationEncounters{}
		err := json.Unmarshal(data, &resp)
		if err != nil {
			return RespLocationEncounters{}, err
		}
		return resp, nil
	}
}
