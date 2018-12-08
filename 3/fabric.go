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

func Overlap(squares []Square) int {
	var grid [1001][1001][]int

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
	//for _, row := range grid {
	//fmt.Println(row)
	//}

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

	fmt.Println("Count: ", Overlap(squares))
}
