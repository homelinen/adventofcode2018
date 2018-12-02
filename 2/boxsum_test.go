package main

import (
	"testing"
)

func TestExactlyTwiceOrThrice(t *testing.T) {

	test_set := []struct {
		id         string
		exp_twice  bool
		exp_thrice bool
	}{
		{"bababc", true, true},
		{"abbcde", true, false},
		{"abcccd", false, true},
		{"aabcdd", true, false},
		{"ababab", false, true},
		{"aaaaba", false, false},
	}

	for _, testrow := range test_set {
		res_twice, res_thrice := ExactlyTwiceOrThrice(testrow.id)
		exp_twice := testrow.exp_twice
		exp_thrice := testrow.exp_thrice

		if res_twice != exp_twice {
			t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", res_twice, exp_twice)
		}

		if res_thrice != exp_thrice {
			t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", res_thrice, exp_thrice)
		}
	}
}
func TestChecksum(t *testing.T) {
	test_set := []struct {
		ids      []string
		expected int
	}{
		{[]string{"bababc"}, 1},
		{[]string{
			"bababc",
			"abbcde",
		}, 2},
		{[]string{
			"bababc",
			"abbcde",
			"abbbde",
		}, 4},
	}

	for _, testrow := range test_set {
		result := Checksum(testrow.ids)
		expected := testrow.expected

		if result != expected {
			t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", result, expected)
		}
	}
}

func TestWordDistance(t *testing.T) {
	test_set := []struct {
		words    []string
		expected int
	}{
		{[]string{
			"fghij",
			"fguij",
		},
			1,
		},
		{[]string{
			"abcde",
			"axcye",
		},
			2,
		},
	}

	for _, testrow := range test_set {
		result := WordDistance(testrow.words[0], testrow.words[1])
		expected := testrow.expected

		if result != expected {
			t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", result, expected)
		}
	}
}

func TestWordSimilarities(t *testing.T) {
	test_set := []struct {
		words    []string
		expected string
	}{
		{[]string{
			"fghij",
			"fguij",
		},
			"fgij",
		},
		{[]string{
			"abcde",
			"axcye",
		},
			"ace",
		},
	}

	for _, testrow := range test_set {
		result := WordSimilarities(testrow.words[0], testrow.words[1])
		expected := testrow.expected

		if result != expected {
			t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", result, expected)
		}
	}
}

func TestCreatePairs(t *testing.T) {
	test_set := []struct {
		words []string
		pairs []Pair
	}{
		{
			[]string{
				"abcde",
				"fghij",
				"klmno",
			},
			[]Pair{
				Pair{
					word_a: "abcde",
					word_b: "fghij",
				},
				Pair{
					word_a: "abcde",
					word_b: "klmno",
				},
				Pair{
					word_a: "fghij",
					word_b: "klmno",
				},
			},
		},
	}

	for _, testrow := range test_set {
		result := CreatePairs(testrow.words)
		expected := testrow.pairs

		if len(result) != len(expected) {
			t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", result, expected)
		}

		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", result[i], expected[i])
			}
		}
	}
}
