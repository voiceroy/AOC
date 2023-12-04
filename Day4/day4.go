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

func generateRange(start, end int) []int {
	var numRange = make([]int, 0, end-start+1)
	for i := start + 1; i <= end; i++ {
		numRange = append(numRange, i)
	}

	return numRange
}

func part2(data []string) int {
	var cardWinsNoOfCards = make(map[int]int)
	var cardQueue []int

	// Calculate what card wins how many cards
	for i, card := range data {
		var currentCardWins int

		numbersPresent := strings.Split(strings.Split(card, ":")[1], " | ")
		winningNumbers, myNumbers := strings.Fields(numbersPresent[0]), strings.Fields(numbersPresent[1])
		for _, number := range myNumbers {
			if slices.Contains(winningNumbers, number) {
				currentCardWins += 1
			}
		}

		cardWinsNoOfCards[i+1] = currentCardWins
		cardQueue = append(cardQueue, i+1)
	}

	var cardsProcessed int
	var queueLength = len(cardQueue)
	for cardsProcessed < queueLength {
		generatedRange := generateRange(cardQueue[cardsProcessed], cardQueue[cardsProcessed]+cardWinsNoOfCards[cardQueue[cardsProcessed]])
		cardQueue = append(cardQueue, generatedRange...)
		cardsProcessed++
		queueLength += len(generatedRange)
	}

	return cardsProcessed
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

	// Part 2
	fmt.Printf("Part 2: %d\n", part2(fileArray))
}
