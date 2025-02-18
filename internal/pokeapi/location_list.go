package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/keyinora/pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(10 * time.Second)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Step 1: Check cache
	if data, found := cache.Get(url); found {
		// If cached data is found, unmarshal it and return
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}
	// Step 2: Network request because data is not cached
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Step 3: Store fetched data in cache
	cache.Add(url, dat)

	// Unmarshal fetched data
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
