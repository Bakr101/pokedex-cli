package main

import "testing"

func TestCleanInput(t *testing.T){
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "BIGBulBA RAiChU FOXfire",
			expected: []string{"bigbulba", "raichu", "foxfire"},
		},
	}

	for _, c := range cases{
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected){
			t.Errorf("words are not complete, length of both slices are not equal")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord{
				t.Errorf("words don't match actual:%v - expected:%v", word, expectedWord)
			}
		}
	}


}