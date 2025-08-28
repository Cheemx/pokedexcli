// internal manages all PokeAPI interactions.

package internal

import (
	"encoding/json"
	"io"
	"net/http"
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

func GetPokeLocationAreas(url string) (PokeResponse, error) {
	res, err := http.Get(url)
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
	return response, nil
}
