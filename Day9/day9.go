package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInts(data []string) [][]int {
	var parsedInts [][]int
	for i, row := range data {
		parsedInts = append(parsedInts, []int{})
		for _, num := range strings.Fields(row) {
			value, _ := strconv.Atoi(num)
			parsedInts[i] = append(parsedInts[i], value)
		}
	}

	return parsedInts
}

func checkZeroes(row []int) bool {
	for _, element := range row {
		if element != 0 {
			return false
		}
	}

	return true
}

func getRowDiff(row []int) []int {
	var newRow []int
	for i := 0; i < len(row)-1; i++ {
		newRow = append(newRow, row[i+1]-row[i])
	}

	return newRow
}

func calculateNextTerm(row []int) int {
	lastElement := row[len(row)-1]
	if checkZeroes(row) {
		return lastElement
	}

	return calculateNextTerm(getRowDiff(row)) + lastElement
}

func partOne(data [][]int) int {
	var sumExtrapolated int
	for _, row := range data {
		sumExtrapolated += calculateNextTerm(row)
	}

	return sumExtrapolated
}

func reverseArray(row []int) []int {
	var reversedArray []int
	for i := range row {
		reversedArray = append(reversedArray, row[len(row)-i-1])
	}

	return reversedArray
}

func partTwo(data [][]int) int {
	var sumExtrapolated int
	for _, row := range data {
		sumExtrapolated += calculateNextTerm(reverseArray(row))
	}

	return sumExtrapolated
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := strings.Split(strings.TrimSpace(string(file)), "\n")
	parsedInts := parseInts(fileArray)

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(parsedInts))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(parsedInts))
}
