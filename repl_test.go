package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU Squirtle",
			expected: []string{"charmander", "bulbasaur", "pikachu", "squirtle"},
		},
		{
			input: " \n  MESSED UP    voLtROn392 \n ",
			expected: []string{"messed", "up", "voltron392"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected length %d but got %d", c.input, len(c.expected), len(actual))
			continue
		}

		for i:= range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input '%s', expected word '%s' but got '%s'", c.input, expectedWord, word)
			}
		}
	}
}
