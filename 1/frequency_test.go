package main

import (
	"testing"
)

// Strings containing addition

// Strings containing subtraction

// Combined samle
func TestFrequency(t *testing.T) {

	sample := []string{
		"-1",
		"-2",
		"-3",
	}

	var result int64
	result = -6

	output := Frequency(sample)
	if output != result {
		t.Errorf("Parsing of sample is incorrect, got: %d, want: %d.", output, result)
	}

}
