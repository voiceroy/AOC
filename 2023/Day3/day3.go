package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Number struct {
	value       int
	row         int
	columnStart int
	columnEnd   int
}

type Gear struct {
	row    int
	column int
}

func checkNumber(data []string, number Number, symbolSet string) bool {
	for cr := max(0, number.row-1); cr <= min(len(data)-1, number.row+1); cr++ {
		for cc := max(0, number.columnStart-1); cc <= min(len((data)[number.row])-1, number.columnEnd); cc++ {
			if !strings.Contains(symbolSet, string((data)[cr][cc])) {
				return true
			}
		}
	}

	return false
}

func getNumbers(data []string) []Number {
	var numbers []Number

	pattern := regexp.MustCompile(`\d+`)
	for i, row := range data {
		matchedNumbers := pattern.FindAllStringIndex(row, -1)
		for _, number := range matchedNumbers {
			value, _ := strconv.Atoi(data[i][number[0]:number[1]])
			numbers = append(numbers, Number{value, i, number[0], number[1]})
		}
	}

	return numbers
}

func partOne(data []string, numbers []Number) int {
	var sumParts int

	for _, number := range numbers {
		if checkNumber(data, number, ".1234567890") {
			sumParts += number.value
		}
	}

	return sumParts
}

func getGears(data []string) map[Gear][]Number {
	var gears = make(map[Gear][]Number)
	for i, row := range data {
		for j, col := range row {
			if string(col) == "*" {
				gears[Gear{i, j}] = nil
			}
		}
	}

	return gears
}

func partTwo(data []string, numbers []Number, gears map[Gear][]Number) int {
	var sumGearRatios int

	for gear := range gears {
		for _, number := range numbers {
			if (number.row-1 <= gear.row && gear.row <= number.row+1) && (number.columnStart-1 <= gear.column && gear.column <= number.columnEnd) {
				gears[gear] = append(gears[gear], number)
			}
		}
	}

	for _, gearNumbers := range gears {
		if len(gearNumbers) == 2 {
			sumGearRatios += gearNumbers[0].value * gearNumbers[1].value
		}
	}

	return sumGearRatios
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open file 'input'\n")
	}

	fileArray := strings.Split(strings.TrimSpace(string(file)), "\n")
	gears := getGears(fileArray)
	numbers := getNumbers(fileArray)

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(fileArray, numbers))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(fileArray, numbers, gears))
}
