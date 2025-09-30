package main

import (
	"time"

	"github.com/jifpbj/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*10)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
