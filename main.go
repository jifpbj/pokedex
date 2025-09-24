package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {

	
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		fmt.Println("Your command was:", words[0])
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

