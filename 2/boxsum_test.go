package main

import (
	"testing"
)

func TestExactlyTwo(t *testing.T) {

	output := ExactlyTwo("bababc")
	result := true
	if output != result {
		t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", output, result)
	}
}
func TestExactlyThree(t *testing.T) {

	output := ExactlyThree("bababc")
	result := true
	if output != result {
		t.Errorf("Parsing of sample is incorrect, got: %v, want: %v.", output, result)
	}
}
