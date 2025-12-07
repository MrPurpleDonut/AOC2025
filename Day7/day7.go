package main

import (
	"fmt"
	"os"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func main() {
	start := time.Now()
	board, err := aoc.MakeMatrix(os.Args[1])
	aoc.HandleError(err)
	count := 0

	counts := make([][]int, len(board))
	for i := range len(board) {
		counts[i] = make([]int, len(board[0]))
	}
	for i, v := range board[0] {
		if v == 'S' {
			counts[0][i] = 1
		}
	}

	for i := range len(board) - 1 {
		count += iterate(board, i+1, counts)
	}

	fmt.Println(count)
	total := 0
	for _, v := range counts[len(board)-1] {
		total += v
	}
	fmt.Println(total)
	fmt.Println(time.Since(start))
}

func iterate(board [][]rune, row int, counts [][]int) int {
	split := 0
	for i, v := range board[row-1] {
		if v == 'S' {
			if board[row][i] == '^' {
				board[row][i+1] = 'S'
				board[row][i-1] = 'S'
				counts[row][i+1] += counts[row-1][i]
				counts[row][i-1] += counts[row-1][i]
				split++
			} else {
				board[row][i] = 'S'
				counts[row][i] += counts[row-1][i]
			}
		}
	}

	return split
}
