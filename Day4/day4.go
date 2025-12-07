package main

import (
	"fmt"
	"os"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

type point struct {
	x int
	y int
}

func main() {
	start := time.Now()

	board, err := aoc.MakeMatrix(os.Args[1])
	aoc.HandleError(err)

	paper := make(map[point]int)

	for x, row := range board {
		for y, val := range row {
			if val == '@' {
				paper[point{x, y}] = 1
			}
		}
	}

	count := 0

	for k := range paper {
		num := 0
		for x := range 3 {
			for y := range 3 {
				if x == 1 && y == 1 {
					continue
				}
				num += paper[point{k.x + x - 1, k.y + y - 1}]
			}
		}
		if num < 4 {
			count++
		}

	}

	fmt.Println(count)

	count2 := 0

	for {
		p, v := removal(&paper)
		if v == 0 {
			break
		}
		count2 += v

		for _, val := range *p {
			paper[val] = 0
		}
	}
	fmt.Println(count2)
	fmt.Println(time.Since(start))
}

func removal(paper *map[point]int) (*[]point, int) {

	count := 0
	removers := make([]point, 0)

	for k := range *paper {
		if (*paper)[k] == 0 {
			continue
		}
		num := 0
		for x := range 3 {
			for y := range 3 {
				if x == 1 && y == 1 {
					continue
				}
				num += (*paper)[point{k.x + x - 1, k.y + y - 1}]
			}
		}
		if num < 4 {
			removers = append(removers, k)
			count++
		}

	}

	return &removers, count
}
