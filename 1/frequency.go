package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Frequency(changes []string) int64 {

	var freq int64
	freq = 0
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

	fmt.Println(Frequency(changes))
}
