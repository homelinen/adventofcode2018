package main

import (
	"fmt"
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
	fmt.Println("vim-go")
}
