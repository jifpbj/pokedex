package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	val, ok := cfg.pokedex[pokemonName]
	if ok {
		println("You already caught", val.Name, "!")
		return nil
	}
	if !ok {

		pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
		if err != nil {
			return errors.New("failed to fetch pokemon: " + err.Error())
		}
		playerChance := rand.Intn(200)
		pokemonChance := pokemon.BaseExperience

		if playerChance > pokemonChance {

			msg := fmt.Sprintf("You caught %s!", pokemon.Name)
			fmt.Println(msg)
			cfg.pokedex[pokemon.Name] = pokemon
		} else {
			fmt.Println("Player chance:", playerChance)
			fmt.Println("Pokemon chance:", pokemonChance)
			fmt.Println("The", pokemon.Name, "escaped!")
		}
	}
	return nil
}
