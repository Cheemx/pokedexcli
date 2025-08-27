package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	mep := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !sc.Scan() {
			break
		}
		cmd := sc.Text()
		val, ok := mep[cmd]
		if !ok {
			fmt.Println("Unknown command")
		}
		switch val.name {
		case "help":
			val.callback()
			for k, v := range mep {
				fmt.Printf("%s: %s\n", k, v.description)
			}
		case "exit":
			val.callback()
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	return nil
}
