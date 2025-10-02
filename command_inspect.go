package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("please provide the name of the Pokémon to inspect")
	}
	caughtPokemon, exists := cfg.caughtPokemon[args[0]]
	if !exists {
		return errors.New("you don't have this Pokémon caught")
	}
	fmt.Printf("Name: %s\n", caughtPokemon.Name)
	fmt.Printf("Height: %d\n", caughtPokemon.Height)
	fmt.Printf("Weight: %d\n", caughtPokemon.Weight)
	fmt.Println("Stats:")

	for _, statValue := range caughtPokemon.Stats {
		fmt.Printf("  -%s: %d\n", statValue.Stat.Name, statValue.BaseStat+statValue.Effort)
	}
	fmt.Println("Types:")
	for _, typeValue := range caughtPokemon.Types {
		fmt.Printf(" - %s\n", typeValue.Type.Name)
	}
	return nil
}
