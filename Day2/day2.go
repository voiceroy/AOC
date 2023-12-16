package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne(MAXDICES []int, games []string) int {
	var validGameIDSum int
	for gameID, game := range games {
		diceSeenThisGame := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		dieCount, dieColor := 0, ""

		rounds := strings.ReplaceAll(strings.TrimSpace(strings.Split(game, ": ")[1]), ";", ",")
		for _, die := range strings.Split(rounds, ",") {
			_, _ = fmt.Sscanf(die, "%d %s", &dieCount, &dieColor)
			diceSeenThisGame[dieColor] = max(diceSeenThisGame[dieColor], dieCount)
		}

		if diceSeenThisGame["red"] <= 12 && diceSeenThisGame["green"] <= 13 && diceSeenThisGame["blue"] <= 14 {
			validGameIDSum += gameID + 1
		}
	}

	return validGameIDSum
}

func partTwo(games []string) int {
	var sumOfProducts int
	for _, game := range games {
		maxDiceThisGame := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		dieCount, dieColor := 0, ""

		rounds := strings.ReplaceAll(strings.TrimSpace(strings.Split(game, ": ")[1]), ";", ",")
		for _, die := range strings.Split(rounds, ",") {
			_, _ = fmt.Sscanf(die, "%d %s", &dieCount, &dieColor)
			maxDiceThisGame[dieColor] = max(maxDiceThisGame[dieColor], dieCount)
		}

		sumOfProducts += maxDiceThisGame["red"] * maxDiceThisGame["green"] * maxDiceThisGame["blue"]
	}

	return sumOfProducts
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open file 'input'")
		return
	}

	MAXDICES := []int{12, 13, 14}
	fileArray := strings.Split(strings.TrimSpace(string(file)), "\n")

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(MAXDICES, fileArray))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(fileArray))
}
