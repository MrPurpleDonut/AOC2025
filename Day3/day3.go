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
	count := 0
	for _, line := range lines {
		count += biggest(line)
	}
	fmt.Println(count)

	count2, count3 := 0, 0
	for _, line := range lines {
		nums := convertToInts(line)
		count2 += biggest2(nums, 2)
		count3 += biggest2(nums, 12)
	}
	fmt.Println(count2)
	fmt.Println(count3)
	fmt.Println(time.Since(start))
}

func convertToInts(line string) []int {
	nums := make([]int, len(line))

	for i, v := range line {
		num, _ := strconv.Atoi(string(v))
		nums[i] = num
	}

	return nums
}

func biggest(line string) int {
	nums := convertToInts(line)
	firstDigit, firstIndex := 0, -1
	for i := 0; i < len(nums)-1; i++ {
		num := nums[i]
		if num > firstDigit {
			firstDigit = num
			firstIndex = i
		}
	}
	secondDigit := 0
	for i := firstIndex + 1; i < len(nums); i++ {
		num := nums[i]
		if num > secondDigit {
			secondDigit = num
		}
	}

	return 10*firstDigit + secondDigit

}

func biggest2(nums []int, length int) int {
	if length < 1 {
		return 0
	}
	firstDigit, firstIndex := -1, 0
	for i := 0; i < len(nums)-length+1; i++ {
		num := nums[i]
		if num > firstDigit {
			firstDigit = num
			firstIndex = i
		}
	}
	rest := biggest2(nums[firstIndex+1:], length-1)
	for range length - 1 {
		firstDigit *= 10
	}
	return rest + firstDigit

}
