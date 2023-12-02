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
	var validGameIdSum int

gameLoop:
	for i, line := range games {
		diceSeenThisRound := []int{0, 0, 0} // [red, green, blue]
		dieCount := 0
		dieColor := ""

		rounds := strings.Split(strings.Split(line, ":")[1], ";")
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
		validGameIdSum += i + 1
	}

	return validGameIdSum
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
}
