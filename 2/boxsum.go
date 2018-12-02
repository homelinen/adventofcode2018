package main

import "bufio"
import "fmt"
import "os"

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

func Checksum(lines []string) int {

	c_twice := 0
	c_thrice := 0
	for _, item := range lines {

		tw, th := ExactlyTwiceOrThrice(item)
		if tw {
			c_twice += 1
		}
		if th {
			c_thrice += 1
		}
	}
	return c_twice * c_thrice
}

func main() {
	file := "input.txt"

	f, _ := os.Open(file)
	reader := bufio.NewReader(f)

	scanner := bufio.NewScanner(reader)

	var changes []string
	for scanner.Scan() {
		changes = append(changes, scanner.Text())
	}
	fmt.Println("Count: ", Checksum(changes))
}
