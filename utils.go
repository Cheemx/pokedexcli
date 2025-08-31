package main

import (
	"math/rand"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func canCatchPokemon(baseExperience int) bool {
	maxExp := 600.0
	normalizedExp := float64(baseExperience) / maxExp

	catchRate := 0.95 - (normalizedExp * 0.9)

	if catchRate < 0.05 {
		catchRate = 0.05
	}
	if catchRate > 0.95 {
		catchRate = 0.95
	}

	randomValue := rand.Float64()
	return randomValue < catchRate
}
