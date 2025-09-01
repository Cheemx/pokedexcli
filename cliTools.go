package main

import (
	"fmt"
	"os"

	"github.com/Cheemx/pokedexcli/internal/pokeapi"
)

const (
	locationAreaAPI = "https://pokeapi.co/api/v2/location-area/"
)

type config struct {
	Previous string
	Next     string
	Client   pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	return nil
}

func commandMap(c *config) error {
	var res pokeapi.PokeResponse
	var err error
	if c.Next == "" {
		res, err = c.Client.GetPokeLocationAreas(locationAreaAPI)
	} else {
		res, err = c.Client.GetPokeLocationAreas(c.Next)
	}
	if err != nil {
		return err
	}

	c.Next = res.Next
	c.Previous = res.Previous
	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(c *config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page.")
		return nil
	}
	var res pokeapi.PokeResponse
	var err error
	res, err = c.Client.GetPokeLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	c.Next = res.Next
	c.Previous = res.Previous
	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(c *config) error {
	if c.Next == "" {
		fmt.Println("No Location Provided")
		return nil
	}

	res, err := c.Client.GetPokeNamesFromLocationAreas(c.Next)
	if err != nil {
		return err
	}
	pokemons := res.PokemonEncounters
	for _, pokeName := range pokemons {
		fmt.Println(pokeName.Pokemon.Name)
	}

	return nil
}

func commandCatch(c *config) error {
	pokemonName := cleanInput(c.Next)
	if pokemonName[0] == "" {
		fmt.Println("No Name Provided")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName[0])

	res, err := c.Client.GetPokemonInformation(pokemonName[0])
	if err != nil {
		return fmt.Errorf("error here: %s", err)
	}

	canCatch := canCatchPokemon(res.BaseExperience)
	if canCatch {
		fmt.Printf("%s was caught!\n", pokemonName[0])
		catchedPokemons[pokemonName[0]] = res
	}
	if !canCatch {
		fmt.Printf("%s escaped!\n", pokemonName[0])
	}
	return nil
}

func commandInspect(c *config) error {
	pokemon, ok := catchedPokemons[c.Next]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.Effort)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("\t- %s\n", typ.Type.Name)
	}

	return nil
}
