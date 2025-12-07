package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

type interval struct {
	start int
	end   int
}

func (i *interval) contains(num int) bool {
	return num >= i.start && num <= i.end
}

func contained(num int, intervals *[]interval) bool {
	for _, i := range *intervals {
		if i.contains(num) {
			return true
		}
	}

	return false
}

func less(a, b interval) int {
	if a.start != b.start {
		return a.start - b.start
	}
	return a.end - b.end
}

func main() {
	start := time.Now()
	rows, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)

	ranges := make([]interval, 0)
	for _, row := range rows {
		row := strings.Replace(row, "-", ",", 1)
		nums, err := aoc.ParseAllInts(row)
		aoc.HandleError(err)
		ranges = append(ranges, interval{nums[0], nums[1]})
	}

	rows2, err := aoc.MakeRows(os.Args[2])
	aoc.HandleError(err)
	count := 0

	for _, row := range rows2 {
		num, err := strconv.Atoi(row)
		aoc.HandleError(err)
		if contained(num, &ranges) {
			count++
		}
	}
	fmt.Println(count)

	slices.SortFunc(ranges,
		func(a, b interval) int {
			if a.start != b.start {
				return a.start - b.start
			}
			return a.end - b.end
		})

	prev, count2 := 0, 0
	for _, i := range ranges {
		if i.end <= prev {
			continue
		}
		if i.start <= prev {
			count2 += i.end - prev
		} else {
			count2 += i.end - i.start + 1
		}

		prev = i.end
	}
	fmt.Println(count2)
	fmt.Println(time.Since(start))
}
