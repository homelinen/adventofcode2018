package main

import "bufio"
import "fmt"
import "os"
import "regexp"
import "strconv"

type Square struct {
	id, x, y, h, w int
}

func Process(line string) Square {
	regex := `^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`

	re := regexp.MustCompile(regex)
	matches := re.FindStringSubmatch(line)

	id, _ := strconv.ParseInt(matches[1], 10, 64)
	x, _ := strconv.ParseInt(matches[2], 10, 64)
	y, _ := strconv.ParseInt(matches[3], 10, 64)
	h, _ := strconv.ParseInt(matches[4], 10, 64)
	w, _ := strconv.ParseInt(matches[5], 10, 64)
	square := Square{
		int(id),
		int(x),
		int(y),
		int(w),
		int(h),
	}

	return square
}

func BuildGrid(squares []Square) [][][]int {
	// Setup grid
	var grid [][][]int = make([][][]int, 1001)

	for i := 0; i < 1001; i++ {
		grid[i] = make([][]int, 1001)
		for j := 0; j < 1001; j++ {
			grid[i][j] = []int{}
		}
	}

	for _, sq := range squares {
		for i := 0; i < sq.h; i++ {
			for j := 0; j < sq.w; j++ {

				if sq.x+j < 1001 && sq.y+i < 1001 {
					grid_sq := grid[sq.x+j][sq.y+i]
					grid[sq.x+j][sq.y+i] = append(grid_sq, sq.id)
				}
			}
		}
	}
	return grid
}

func Overlap(grid [][][]int) int {
	overlaps := 0

	for _, x := range grid {
		for _, y := range x {
			if len(y) > 1 {
				overlaps += 1
			}
		}
	}

	return overlaps
}

func Untouched(grid [][][]int, squares []Square) int {
	untouched := make(map[int]int, 0)

	for _, x := range grid {
		for _, y := range x {
			if len(y) > 1 {
				for _, sq := range y {
					untouched[sq] += 1
				}
			}
		}
	}

	the_untouched := 0
	for _, s := range squares {
		found := false
		for k, _ := range untouched {
			if s.id == k {
				found = true
			}
		}

		if !found {
			the_untouched = s.id
			break
		}
	}

	return the_untouched
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

	squares := make([]Square, 0)
	for _, line := range changes {
		// Clean line`
		squares = append(squares, Process(line))
	}

	grid := BuildGrid(squares)
	fmt.Println("Count: ", Overlap(grid))

	fmt.Println("Untouched: ", Untouched(grid, squares))
}
