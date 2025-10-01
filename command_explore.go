package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a location name")
	}
	deepLocation, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}
	Encounters := deepLocation.PokemonEncounters

	for _, encounter := range Encounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
