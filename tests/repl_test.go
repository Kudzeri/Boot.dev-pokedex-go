package tests

import (
	"github.com/Kudzeri/Boot.dev-pokedex-go/repl"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "GOODBYE world",
			expected: []string{
				"goodbye",
				"world",
			},
		},
		{
			input: "123 33",
			expected: []string{
				"123",
				"33",
			},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, cs := range cases {
		actual := repl.CleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v",
				len(actual),
				len(cs.expected),
			)
			continue
		}
		for i, _ := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v",
					actualWord,
					expectedWord,
				)
			}
		}
	}
}
