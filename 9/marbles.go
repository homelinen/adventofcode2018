package main

import "fmt"
import "container/ring"

func MaxSlice(sl []int) (pos int, max int) {
	max = sl[0]
	pos = 0
	for i, v := range sl {

		if v > max {
			max = v
			pos = i
		}
	}
	return pos, max
}

func GameOfMarbles(players int, final_marble int) int {
	player_score := make([]int, players)

	board := ring.New(1)
	r := board.Next()
	r.Value = 0

	cur_player := 0
	for i := 1; i <= final_marble; i++ {

		if i%23 == 0 {
			board = board.Move(board.Len() - 7)
			removed := board.Next().Value
			board.Unlink(1)
			player_score[cur_player%players] += i
			player_score[cur_player%players] += removed.(int)

			cur_player += 1
			continue
		}

		board = board.Move(2)
		new_ring := ring.New(1)
		nr := new_ring.Next()
		nr.Value = i

		board.Link(new_ring)
		cur_player += 1
	}
	// Iterate through the ring and print its contents
	//board.Do(func(p interface{}) {
	//fmt.Println(p.(int))
	//})

	player_score[0] = 1

	_, score := MaxSlice(player_score)

	return score
}

func main() {
	fmt.Println("Hi")
}
