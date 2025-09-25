package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer res.Body.Close()

	var locations Location
	decoder := json.NewDecoder(res.Body)
	err2 := decoder.Decode(&locations)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return nil
}
