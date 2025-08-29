// internal manages all PokeAPI interactions.

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Cheemx/pokedexcli/internal/pokecache"
)

// internal is a special directory name recognised by the go tool
// which will prevent one package from being imported by another
// unless both share a common ancestor.

type PokeResponse struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"results,omitempty"`
}

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetPokeLocationAreas(url string) (PokeResponse, error) {
	if val, ok := c.cache.Get(url); ok {
		var response PokeResponse
		err := json.Unmarshal(val, &response)
		if err != nil {
			return PokeResponse{}, err
		}
		return response, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return PokeResponse{}, err
	}
	var response PokeResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return PokeResponse{}, err
	}

	c.cache.Add(url, body)
	return response, nil
}
