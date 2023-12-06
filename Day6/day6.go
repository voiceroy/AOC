package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInts(array []string) []int {
	var parsed []int
	for _, num := range array {
		value, _ := strconv.Atoi(num)
		parsed = append(parsed, value)
	}

	return parsed
}

func part1(data []string) int {
	var product = 1

	timeArray := parseInts(strings.Fields(strings.Split(data[0], ":")[1]))
	distanceArray := parseInts(strings.Fields(strings.Split(data[1], ":")[1]))
	for i, time := range timeArray {
		distance := distanceArray[i]
		possible := 0

		for j := 0; j < time+1; j++ {
			if j*(time-j) > distance {
				possible++
			}
		}

		product *= possible
	}

	return product
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := strings.Split(string(file), "\n")
	fileArray = fileArray[:len(fileArray)-1]

	// Part 1
	fmt.Printf("Part 1: %d\n", part1(fileArray))
}
