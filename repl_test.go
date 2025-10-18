package main

import (
	"testing"
)

func TestCleanInput (t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world USA France Vietnam!",
			expected: []string{"hello", "world", "usa", "france", "vietnam!"},
		},
		{
			input: "   hello   world   USA   !   ",
			expected: []string{"hello", "world", "usa", "!"},
		},
		{
			input: "    ",
			expected: []string{},
		},
	}
	
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(`
			Lengths don't match
			Expected: %v
			Actual: %v
			`, c.expected, actual)
			continue
		}
		
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf(`
				Mismatch found!
				Expected: %s
				Actual: %s
				`, c.expected[i], actual[i])
			}
		}
	}
}