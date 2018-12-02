package main

import "fmt"

func ExactlyTwiceOrThrice(boxid string) (bool, bool) {

	var twice, thrice bool
	letter_counts := make(map[rune]int)

	for _, letter := range boxid {
		if letter_counts[letter] != 0 {
			letter_counts[letter] += 1
		} else {
			letter_counts[letter] = 1
		}
	}

	for _, v := range letter_counts {
		if v == 2 {
			twice = true
		} else if v == 3 {
			thrice = true
		}
	}

	return twice, thrice
}

func ExactlyThree(boxid string) bool {
	return false
}

func main() {
	fmt.Println("vim-go")
}
