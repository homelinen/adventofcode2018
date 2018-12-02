package main

import (
	"testing"
)

func TestFrequency(t *testing.T) {

	sample := []string{
		"-1",
		"-2",
		"-3",
	}

	var result int64
	result = -6

	output := Frequency(sample, 0)
	if output != result {
		t.Errorf("Parsing of sample is incorrect, got: %d, want: %d.", output, result)
	}

	// Example 2

	sample = []string{
		"+1",
		"+1",
		"-2",
	}

	result = 0

	output = Frequency(sample, 0)
	if output != result {
		t.Errorf("Parsing of sample is incorrect, got: %d, want: %d.", output, result)
	}
}

func TestFrequencyTwice(t *testing.T) {

	sample := []string{
		"+3",
		"+3",
		"+4",
		"-2",
		"-4",
	}

	var result int64
	result = 10

	output := FrequencyTwice(sample)
	if output != result {
		t.Errorf("Parsing of sample is incorrect, got: %d, want: %d.", output, result)
	}
}
