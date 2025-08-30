package pokeapi

import (
	"net/http"
	"time"

	"github.com/Cheemx/pokedexcli/internal/pokecache"
)

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

type LocationArea struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

// Encounter method related structures
type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodRates struct {
	EncounterMethod         EncounterMethod           `json:"encounter_method"`
	EncounterVersionDetails []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}

// Location structure
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Version structure
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Language and names structures
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Names struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}

// Pokemon encounter structures
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterDetails struct {
	Chance          int    `json:"chance"`
	ConditionValues []any  `json:"condition_values"`
	MaxLevel        int    `json:"max_level"`
	Method          Method `json:"method"`
	MinLevel        int    `json:"min_level"`
}

type PokemonVersionDetails struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Version          Version            `json:"version"`
}

type PokemonEncounters struct {
	Pokemon        Pokemon                 `json:"pokemon"`
	VersionDetails []PokemonVersionDetails `json:"version_details"`
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
