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

	rows, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)

	operationString := rows[len(rows)-1]
	totals, err := aoc.ParseAllInts(rows[0])
	aoc.HandleError(err)

	numStrings := rows[1 : len(rows)-1]

	ops := parseOperations(operationString)

	for _, row := range numStrings {
		vals, err := aoc.ParseAllInts(row)
		aoc.HandleError(err)
		for i, v := range vals {
			if (*ops)[i] == '+' {
				totals[i] += v
			} else {
				totals[i] *= v
			}
		}

	}
	count := 0
	for _, v := range totals {
		count += v
	}
	fmt.Println(count)
	fmt.Println(part2(rows))
	fmt.Println(time.Since(start))
}

func parseOperations(s string) *[]rune {
	operations := make([]rune, 0)

	for _, v := range s {
		if v != ' ' {
			operations = append(operations, v)
		}
	}

	return &operations
}

func part2(rows []string) int {
	count := 0
	currentOp := ' '
	currentVals := 0
	for i, op := range rows[len(rows)-1] {
		if op != ' ' {
			count += currentVals
			currentOp = op
			currentVals = 0
		}

		currentNum := 0
		for _, row := range rows[:len(rows)-1] {
			if row[i] != ' ' {
				digit, err := strconv.Atoi(string(row[i]))
				aoc.HandleError(err)
				currentNum *= 10
				currentNum += digit
			}
		}
		if currentNum == 0 {
			continue
		}
		if currentVals == 0 {
			currentVals = currentNum
			continue
		}
		if currentOp == '+' {
			currentVals += currentNum
		} else {
			currentVals *= currentNum
		}

	}

	return count + currentVals

}
