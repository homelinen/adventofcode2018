package main

import "testing"

func TestGameOfMarbles(t *testing.T) {

	test_set := []struct {
		players     int
		last_marble int
		expected    int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
		{493, 71863, 367802},
		//{493, 7186300, 367802},
	}
	for _, testrow := range test_set {
		result := GameOfMarbles(testrow.players, testrow.last_marble)

		if result != testrow.expected {
			t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", result, testrow.expected)
		}
	}
}
