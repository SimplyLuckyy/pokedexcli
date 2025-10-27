package main

import (
	"testing"
	"fmt"
)


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input 	 string
		expected []string
	}{
		{
			input:	"   hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:	"   Charmander    BULBASAUR squiRTLE ",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		{
			input:	"     ",
			expected: []string{},
		},
		}
	
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("slice lengths don't match")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {t.Errorf("words don't match")}
		} 
	}
}