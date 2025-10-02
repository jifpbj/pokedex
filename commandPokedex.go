package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("no Pokémon caught yet")
	}
	for _, pokemonName := range cfg.caughtPokemon {
		fmt.Println(" -", pokemonName.Name)
	}
	return nil
}
