package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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

	// Step 2: Network request because data is not cached
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationEncounters{}, err
	}

	// Step 3: Process the Request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationEncounters{}, err
	}

	// close connection once function exit
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationEncounters{}, err
	}
	// Step 4: Store fetched data in cache
	cache.Add(url, dat)

	// Step 5: Unmarshal fetched data
	exploreResp := RespLocationEncounters{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return RespLocationEncounters{}, err
	}
	return exploreResp, nil
}
