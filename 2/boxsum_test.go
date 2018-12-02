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
