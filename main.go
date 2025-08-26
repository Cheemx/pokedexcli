package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !sc.Scan() {
			break
		}
		// cmd := sc.Text()
	}
}
