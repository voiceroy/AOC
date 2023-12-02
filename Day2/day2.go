package main

import (
	"fmt"
	"os"
	"strings"
)

func validRound(MAXDICES []int, game []int) bool {
	for i := 0; i < len(MAXDICES); i++ {
		if !(MAXDICES[i] >= game[i]) {
			return false
		}
	}

	return true
}

func part1(MAXDICES []int, games []string) int {
	var validGameIDSum int

gameLoop:
	for gameID, game := range games {
		diceSeenThisRound := []int{0, 0, 0} // [red, green, blue]
		dieCount := 0
		dieColor := ""

		rounds := strings.Split(strings.Split(game, ":")[1], ";")
		for _, round := range rounds {
			for _, die := range strings.Split(round, ",") {
				_, _ = fmt.Sscanf(die, "%d %s", &dieCount, &dieColor)

				switch dieColor {
				case "red":
					diceSeenThisRound[0] = dieCount
				case "green":
					diceSeenThisRound[1] = dieCount
				case "blue":
					diceSeenThisRound[2] = dieCount
				}

				if !validRound(MAXDICES, diceSeenThisRound) {
					continue gameLoop
				}
			}
		}
		validGameIDSum += gameID + 1
	}

	return validGameIDSum
}

func product(game []int) int {
	result := 1
	for i := 0; i < len(game); i++ {
		result *= game[i]
	}

	return result
}

func part2(games []string) int {
	var sumOfProducts int

	for _, game := range games {
		maxDiceThisRound := []int{0, 0, 0}
		dieCount := 0
		dieColor := ""

		rounds := strings.Split(strings.Split(game, ":")[1], ";")
		for _, round := range rounds {
			for _, die := range strings.Split(round, ",") {
				_, _ = fmt.Sscanf(die, "%d %s", &dieCount, &dieColor)

				switch dieColor {
				case "red":
					maxDiceThisRound[0] = max(maxDiceThisRound[0], dieCount)
				case "green":
					maxDiceThisRound[1] = max(maxDiceThisRound[1], dieCount)
				case "blue":
					maxDiceThisRound[2] = max(maxDiceThisRound[2], dieCount)
				}
			}
		}
		sumOfProducts += product(maxDiceThisRound)
	}

	return sumOfProducts
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open file 'input'")
		return
	}

	var MAXDICES = []int{12, 13, 14}
	fileArray := strings.Split(string(file), "\n")

	// Part 1
	fmt.Printf("Part 1: %d\n", part1(MAXDICES, fileArray[:len(fileArray)-1]))

	// Part 2
	fmt.Printf("Part 2: %d\n", part2(fileArray[:len(fileArray)-1]))
}
