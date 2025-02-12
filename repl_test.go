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
			input:    " hello world ",
			expected: []string{"hello", "world"},
		}, //More spaces around
		{
			input:    " hello  world  ",
			expected: []string{"hello", "world"},
		}, //Big Spaces
		{
			input:    "       hello         world       ",
			expected: []string{"hello", "world"},
		}, //Tabs
		{
			input:    "	hello	world	",
			expected: []string{"hello", "world"},
		}, //Newlines
		{
			input:    " hello\n world \n",
			expected: []string{"hello", "world"},
		}, //Special Characters
		{
			input:    " %$%hello$#\n world!!! \n!!!",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hElLo WORLd ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "		",
			expected: []string{},
		},
		{
			input:    " Hello123 233World23 ",
			expected: []string{"hello", "world"},
		},
		//add more cases
	}

	for test_index, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Test #%v: Actual length %v does not match expected length %v", test_index, len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test #%v: Actual word of %v does not match expected word of %v", test_index, word, expectedWord)
			}
		}
	}
}
