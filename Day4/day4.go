package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func part1(data []string) int {
	var allCardsWorth int

	for _, line := range data {
		var currentCardWorth int

		numbersPresent := strings.Split(strings.Split(line, ":")[1], " | ")
		winningNumbers, myNumbers := strings.Fields(numbersPresent[0]), strings.Fields(numbersPresent[1])
		for _, number := range myNumbers {
			if slices.Contains(winningNumbers, number) {
				currentCardWorth = max(1, currentCardWorth*2)
			}
		}

		allCardsWorth += currentCardWorth
	}

	return allCardsWorth
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := strings.Split(string(file), "\n")

	// Part 1
	fmt.Printf("Part 1: %d\n", part1(fileArray[:len(fileArray)-1]))
}
