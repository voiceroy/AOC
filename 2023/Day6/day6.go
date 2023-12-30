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

func partOne(data []string) int {
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

func partTwo(data []string) int {
	var product = 1

	time, _ := strconv.Atoi(strings.Replace(strings.Split(data[0], ":")[1], " ", "", -1))
	distance, _ := strconv.Atoi(strings.Replace(strings.Split(data[1], ":")[1], " ", "", -1))
	possible := 0
	for i := 0; i < time+1; i++ {
		if i*(time-i) > distance {
			possible++
		}
	}

	product *= possible
	return product
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := strings.Split(strings.TrimSpace(string(file)), "\n")

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(fileArray))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(fileArray))
}
