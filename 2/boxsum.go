package main

import "bufio"
import "fmt"
import "os"
import "sort"

type Pair struct {
	word_a   string
	word_b   string
	distance int
}

func (a Pair) Compare(b Pair) bool {

	if (a.word_a == b.word_a && a.word_b == b.word_b) ||
		(a.word_b == b.word_a && a.word_a == b.word_b) {
		return true
	}
	return false
}

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

// Just difference part of Levenshtein
func WordDistance(a string, b string) int {

	distance := 0

	if len(a) != len(b) {
		fmt.Println("Different Lengths issue")
	}

	for i, _ := range a {
		if a[i] != b[i] {
			distance += 1
		}
	}

	return distance
}

// Just difference part of Levenshtein

func WordSimilarities(a string, b string) string {

	difference := ""

	if len(a) != len(b) {
		fmt.Println("Different Lengths issue")
	}

	for i, _ := range a {
		if a[i] == b[i] {
			difference = difference + string(a[i])
		}
	}

	return difference
}

func Any(vs []Pair, f func(Pair) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func CreatePairs(ids []string) []Pair {
	// FIXME: Thisis super close

	// pairs length should be n^2?
	pairs := make([]Pair, 0)
	for i, _ := range ids {
		for j, _ := range ids {
			if ids[i] == ids[j] {
				continue
			}
			p := Pair{
				word_a: ids[i],
				word_b: ids[j],
			}

			if !Any(pairs, func(s Pair) bool {
				return s.Compare(p)
			}) {
				pairs = append(pairs, p)
			}
		}
	}

	return pairs
}

func GetClosestSharedLetters(ids []string) string {

	pairs := CreatePairs(ids)

	fmt.Println(pairs)
	distances := make([]Pair, 0)

	for _, p := range pairs {
		p.distance = WordDistance(p.word_a, p.word_b)
		distances = append(distances, p)
	}

	sort.Slice(distances, func(i, j int) bool { return distances[i].distance < distances[j].distance })
	fmt.Println(distances[0])

	letters := WordSimilarities(distances[0].word_a, distances[0].word_b)

	return letters
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

	fmt.Println("Similar Closest: ", GetClosestSharedLetters(changes))
}
