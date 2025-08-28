package main

import (
	"fmt"
	"os"

	"github.com/Cheemx/pokedexcli/internal"
)

const (
	locationAreaAPI = "https://pokeapi.co/api/v2/location-area/"
)

type config struct {
	Previous string
	Next     string
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
	var res internal.PokeResponse
	var err error
	if c.Next == "" {
		res, err = internal.GetPokeLocationAreas(locationAreaAPI)
	} else {
		res, err = internal.GetPokeLocationAreas(c.Next)
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
	var res internal.PokeResponse
	var err error
	res, err = internal.GetPokeLocationAreas(c.Previous)
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
