package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type number struct {
	value       string
	row         int
	columnStart int
	columnEnd   int
}

const validSymbols = "+-*/@&$#=%"

func checkNumber(data []string, possibleNumber number) bool {
	var validNumber bool
	columnStart := max(0, possibleNumber.columnStart-1)
	columnEnd := min(possibleNumber.columnEnd+1, len(data[0]))

	// Top Row
	if possibleNumber.row-1 >= 0 {
		validNumber = validNumber || strings.ContainsAny(data[possibleNumber.row-1][columnStart:columnEnd], validSymbols)
	}

	// Bottom Row
	if possibleNumber.row+1 < len(data) {
		validNumber = validNumber || strings.ContainsAny(data[possibleNumber.row+1][columnStart:columnEnd], validSymbols)
	}

	// Current Row
	validNumber = validNumber || strings.ContainsAny(data[possibleNumber.row][columnStart:columnEnd], validSymbols)

	return validNumber
}

func part1(data []string) int {
	var sumPartNumbers int
	var numberLocations []number

	for row, line := range data {
		var currentNum = number{"", row, -1, -1}

		for column, char := range line {
			if unicode.IsDigit(char) {
				if currentNum.columnStart == -1 {
					currentNum.columnStart = column
					currentNum.columnEnd = column
				}

				currentNum.value += string(char)
				currentNum.columnEnd += 1
			} else {
				if currentNum.columnStart != -1 {
					numberLocations = append(numberLocations, currentNum)
				}

				currentNum.value = ""
				currentNum.columnStart = -1
				currentNum.columnEnd = -1
			}
		}

		if currentNum.columnStart != -1 {
			numberLocations = append(numberLocations, currentNum)
		}
	}

	for _, num := range numberLocations {
		if checkNumber(data, num) {
			value, _ := strconv.Atoi(num.value)
			sumPartNumbers += value
		}
	}

	return sumPartNumbers
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open file 'input'\n")
	}

	fileArray := strings.Split(string(file), "\n")
	fileArray = fileArray[:len(fileArray)-1]

	// Part 1
	fmt.Printf("Part 1: %d\n", part1(fileArray))
}
