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

type gear struct {
	row      int
	column   int
	adjacent int
}

func checkNumber(data []string, possibleNumber number, validSymbols string) bool {
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

func getNumberLocations(data []string) []number {
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

	return numberLocations
}

func part1(data []string) int {
	var sumPartNumbers int
	var numberLocations = getNumberLocations(data)

	for _, num := range numberLocations {
		if checkNumber(data, num, "+-*/@&$#=%") {
			value, _ := strconv.Atoi(num.value)
			sumPartNumbers += value
		}
	}

	return sumPartNumbers
}

func getGearLocations(data []string) []gear {
	var gearLocations []gear
	for i, line := range data {
		for j, char := range line {
			if string(char) == "*" {
				gearLocations = append(gearLocations, gear{i, j, 0})
			}
		}
	}

	return gearLocations
}

func part2(data []string) int {
	var sumOfGearRatios int
	var numberLocations []number
	var gearLocations = getGearLocations(data)

	for _, num := range getNumberLocations(data) {
		if checkNumber(data, num, "*") {
			numberLocations = append(numberLocations, num)
		}
	}

	for _, currentGear := range gearLocations {
		nums := make([]number, 0)
		for _, num := range numberLocations {
			if (num.columnStart <= currentGear.column && currentGear.column+1 <= num.columnEnd && max(0, num.row-1) <= currentGear.row && currentGear.row <= min(num.row+1, len(data))) || ((currentGear.column == num.columnStart-1 || currentGear.column == num.columnEnd) && (max(0, num.row-1) <= currentGear.row && currentGear.row <= min(num.row+1, len(data)))) {
				currentGear.adjacent++
				nums = append(nums, num)
			}
		}

		if len(nums) == 2 {
			gearRatio := 1

			for _, num := range nums {
				value, _ := strconv.Atoi(num.value)
				gearRatio *= value
			}
			fmt.Printf("Gear Pairs: %v\n", nums)
			sumOfGearRatios += gearRatio
		}
	}

	return sumOfGearRatios
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

	// Part 2
	fmt.Printf("Part 2: %d\n", part2(fileArray))
}
