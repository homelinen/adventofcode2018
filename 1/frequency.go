package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Frequency(changes []string, start_freq int64) int64 {

	var freq int64
	freq = start_freq
	for _, change := range changes {
		op := change[0]
		factor, _ := strconv.ParseInt(change[1:], 10, 64)

		if op == '+' {
			freq += factor
		} else if op == '-' {
			freq -= factor
		}
	}
	return freq
}

// https://stackoverflow.com/a/39868255
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func Any(vs []int64, f func(int64) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func FrequencyTwice(changes []string) int64 {

	var found_freq []int64
	var first_pass int64

	var found bool

	for _, n := range makeRange(1, len(changes)) {
		first_pass = Frequency(changes[:n], 0)

		found_freq = append(found_freq, first_pass)
	}
	var start_freq int64 = first_pass

	for found == false {
		// FIXME: Deadslow, calculating the pass from scratch each time 2n time
		for _, n := range makeRange(1, len(changes)) {
			first_pass = Frequency(changes[:n], start_freq)
			//fmt.Println(first_pass)

			if Any(found_freq, func(c int64) bool {
				return first_pass == c
			}) {
				//fmt.Println("Found")
				found = true
				break
			}

		}
		start_freq = first_pass
	}

	return first_pass
}

func main() {
	fmt.Println("frequency")

	file := "input.txt"

	f, _ := os.Open(file)
	reader := bufio.NewReader(f)

	scanner := bufio.NewScanner(reader)

	var changes []string
	for scanner.Scan() {
		changes = append(changes, scanner.Text())
	}

	fmt.Println("Part 1: ", Frequency(changes, 0))
	fmt.Println("Part 2: ", FrequencyTwice(changes))
}
