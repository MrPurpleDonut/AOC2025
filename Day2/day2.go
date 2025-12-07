package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func main() {
	start := time.Now()

	data, err := os.ReadFile(os.Args[1])
	aoc.HandleError(err)

	count, count2 := 0, 0

	ranges := strings.Split(string(data), ",")
	for _, endpoints := range ranges {
		endpoints = strings.Replace(endpoints, "-", ",", 2)
		nums, err := aoc.ParseAllInts(endpoints)
		aoc.HandleError(err)
		for i := nums[0]; i <= nums[1]; i++ {
			if isValidNum(i) {
				count += i
			}
			if isValidNum2(i) {
				count2 += i
			}
		}

	}
	fmt.Println(count, count2)

	fmt.Println(time.Since(start))
}

func isValidNum(num int) bool {
	val := strconv.Itoa(num)
	if len(val)%2 == 1 {
		return false
	}
	return val[:len(val)/2] == val[len(val)/2:]
}

func isValidNum2(num int) bool {
	val := strconv.Itoa(num)
	l := len(val)
	for i := range l - 1 {
		i += 2
		if l%i != 0 {
			continue
		}
		good := true
		length := l / i
		for j := range i - 1 {
			if val[j*length:(j+1)*length] != val[(j+1)*length:(j+2)*length] {
				good = false
			}
		}
		if good {
			return true
		}
	}

	return false
}
