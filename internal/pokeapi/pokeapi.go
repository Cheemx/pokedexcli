// internal manages all PokeAPI interactions.

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// internal is a special directory name recognised by the go tool
// which will prevent one package from being imported by another
// unless both share a common ancestor.

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

func (c *Client) GetPokeNamesFromLocationAreas(locationName string) (LocationArea, error) {
	url := pokeapi + locationName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var response LocationArea
	err = json.Unmarshal(body, &response)
	if err != nil {
		return LocationArea{}, err
	}
	return response, nil
}
