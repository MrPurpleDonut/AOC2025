package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func main() {
	start := time.Now()

	lines, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)
	count := 50
	rotations := 0

	neg := 1
	for _, line := range lines {
		if line[0] == 'R' {
			neg = 1
		} else {
			neg = -1
		}
		num, _ := strconv.Atoi(line[1:])
		for range num {
			count += neg
			count %= 100
			//fmt.Println(line, neg, num, count)
			if count == 0 {
				rotations++
			}
		}
	}

	fmt.Println(rotations)
	fmt.Println(time.Since(start))
}
