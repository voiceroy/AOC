package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(data string) int {
	calibrationValue := 0
	for _, line := range strings.Split(data, "\n") {
		leftMost := -1
		rightMost := -1

		for _, char := range line {
			if value, err := strconv.Atoi(string(char)); err == nil {
				if leftMost == -1 {
					leftMost = value
					rightMost = leftMost
				} else {
					rightMost = value
				}
			}
		}
		if leftMost != -1 {
			calibrationValue += leftMost*10 + rightMost
		}
	}

	return calibrationValue
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open file 'input'\n")
		return
	}
	fileString := string(file)

	// Part 1
	fmt.Printf("Part 1: %d\n", part1(fileString))
}
