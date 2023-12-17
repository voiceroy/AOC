package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func partOne(data []string) int {
	var totalPoints int

	for _, line := range data {
		numbers := strings.Split(strings.Split(line, ":")[1], " | ")
		winningNumbers, myNumbers := strings.Fields(numbers[0]), strings.Fields(numbers[1])

		var currentCardPoints int
		for _, number := range myNumbers {
			if slices.Contains(winningNumbers, number) {
				currentCardPoints = max(1, currentCardPoints*2)
			}
		}

		totalPoints += currentCardPoints
	}

	return totalPoints
}

func generateRange(start, end int) []int {
	var numRange = make([]int, 0, end-start+1)
	for i := start + 1; i <= end; i++ {
		numRange = append(numRange, i)
	}

	return numRange
}

func partTwo(data []string) int {
	var cardsWon = make([]int, len(data))
	for i := range cardsWon {
		cardsWon[i] = 1
	}

	for i, card := range data {
		numbersPresent := strings.Split(strings.Split(card, ":")[1], " | ")
		winningNumbers, myNumbers := strings.Fields(numbersPresent[0]), strings.Fields(numbersPresent[1])

		var matches int
		for _, number := range myNumbers {
			if slices.Contains(winningNumbers, number) {
				matches += 1
			}
		}

		for wonCard := i + 1; wonCard < i+matches+1; wonCard++ {
			cardsWon[wonCard] += cardsWon[i]
		}
	}

	cardsProcessed := 0
	for _, won := range cardsWon {
		cardsProcessed += won
	}

	return cardsProcessed
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
