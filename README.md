# PokedexCLI

A command-line Pokédex tool built in Go that allows you to explore the Pokémon world, catch Pokémon, and manage your collection.

## Features

- Explore location areas and discover Pokémon
- Catch and collect Pokémon 
- View details of caught Pokémon
- Thread-safe in-memory cache with TTL for API responses
- Comprehensive test coverage for cache functionality

## Installation

```bash
git clone https://github.com/Cheemx/pokedexcli
cd pokedexcli
go build
```

## Commands

- `help` - Display available commands
- `map` - Show 20 location areas in the Pokémon world
- `mapb` - Show previous 20 locations
- `explore <location>` - List Pokémon found in a specific location
- `catch <pokemon>` - Attempt to catch a Pokémon
- `inspect <pokemon>` - View details of a caught Pokémon
- `pokedex` - List all caught Pokémon
- `exit` - Quit the application

## Usage

```bash
./pokedexcli
Pokedex > help
Pokedex > map
Pokedex > explore canalave-city-area
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > pokedex
```

## Testing

```bash
go test ./...
```