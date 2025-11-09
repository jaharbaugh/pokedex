package main

import(
	"testing"
)



func TestCleanInput(t *testing.T){

	cases := []struct {
		input string
		expected []string
}{
	{
		input: " hello world ",
		expected: []string{"hello", "world"},
	},
	{
		input: "HellO WorLD",
		expected: []string{"hello", "world"},

	},
	{
		input: "Pikachu Charizard Sandslash",
		expected: []string{"pikachu", "charizard", "sandslash"},
	},
}

	for _, c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Failed: Length mismatch: got %d; expected %d", len(actual), len(c.expected))
		} 
		for i:= range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("Failed word mismatch: got %q; expected %q", actual[i], c.expected[i])
			}
		}
	}

}