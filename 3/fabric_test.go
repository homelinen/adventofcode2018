package main

import "testing"

func TestProcessing(t *testing.T) {

	test_set := []struct {
		input      string
		exp_square Square
	}{
		{"#1 @ 1,3: 4x4", Square{1, 1, 3, 4, 4}},
	}

	for _, testrow := range test_set {
		res_square := Process(testrow.input)
		exp_square := testrow.exp_square

		if res_square != exp_square {
			t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", res_square, exp_square)
		}
	}
}

func TestOverlap(t *testing.T) {

	test_set := []struct {
		square1 Square
		square2 Square
		overlap int
	}{
		{
			Square{1, 1, 3, 4, 4},
			Square{2, 3, 1, 4, 4},
			4,
		},
		{
			Square{1, 1, 1, 5, 2},
			Square{2, 2, 2, 2, 2},
			2,
		},
	}

	for _, testrow := range test_set {
		result := Overlap([]Square{testrow.square1, testrow.square2})
		expected := testrow.overlap

		if result != expected {
			t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", result, expected)
		}
	}
}
