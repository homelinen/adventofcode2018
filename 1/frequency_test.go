package main

import (
	"testing"
)

// Strings containing addition

// Strings containing subtraction

// Combined samle
func TestFrequency(t *testing.T) {

	sample := "-1\n-2\n-3"

	result := -6

	output := Frequency(sample)
	if output != result {
		t.Errorf("Parsing of sample is incorrect, got: %d, want: %d.", output, result)
	}

}
