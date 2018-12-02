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

func FrequencyTwice(changes []string) int64 {

	found := false

	var second_pass int64
	num := 1

	// FIXME: Deadslow, calculating the pass from scratch each time 2n time
	for found == false {
		found = false
		first_pass := Frequency(changes[:num], 0)

		fmt.Println("f: ", first_pass)
		for _, n := range makeRange(1, len(changes)) {
			second_pass = Frequency(changes[:n], first_pass)
			fmt.Println("s:", second_pass)
			if second_pass == first_pass {
				fmt.Println("Found")
				found = true
			}
		}

		num += 1
	}

	return second_pass
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
