package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type hand struct {
	hand            string
	morphedHand     string
	handType        int
	morphedHandType int
	bid             int
}

func getHandType(hand string) int {
	var handType int

	cardFreq := make(map[rune]int)
	for _, card := range hand {
		cardFreq[card]++
	}

	var frequencies = make([]int, 0, len(cardFreq))
	for _, freq := range cardFreq {
		frequencies = append(frequencies, freq)
	}

	slices.Sort(frequencies)
	if len(frequencies) == 1 {
		handType = 7
	} else if len(frequencies) == 2 {
		if slices.Equal(frequencies, []int{1, 4}) {
			handType = 6
		} else {
			handType = 5
		}
	} else if len(frequencies) == 3 {
		if slices.Equal(frequencies, []int{1, 1, 3}) {
			handType = 4
		} else {
			handType = 3
		}
	} else if len(frequencies) == 4 {
		handType = 2
	} else if len(frequencies) == len(hand) {
		handType = 1
	}

	return handType
}

func morphHand(oldHand string) string {
	cardFreq := make(map[rune]int)
	for _, card := range oldHand {
		cardFreq[card]++
	}

	maxK := rune(0)
	maxV := 0

	for k, v := range cardFreq {
		if v > maxV && string(k) != "J" {
			maxV = v
			maxK = k
		}
	}

	if maxK == 0 {
		return oldHand
	} else {
		return strings.Replace(oldHand, "J", string(maxK), -1)
	}
}

func processHands(data []string) []hand {
	var processedHands []hand
	var handString string
	var bid int

	for _, line := range data {
		_, _ = fmt.Sscanf(line, "%s %d", &handString, &bid)
		morphedString := morphHand(handString)
		processedHands = append(processedHands, hand{handString, morphedString, getHandType(handString), getHandType(morphedString), bid})
	}

	return processedHands
}

func compareHands(hand1, hand2 hand) int {
	var result int
	var cardPriority map[string]int

	cardPriority = map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}

	if hand1.handType == hand2.handType {
		for i := 0; i < len(hand1.hand); i++ {
			if cardPriority[string(hand1.hand[i])] > cardPriority[string(hand2.hand[i])] {
				result = 1
				break
			} else if cardPriority[string(hand1.hand[i])] < cardPriority[string(hand2.hand[i])] {
				result = -1
				break
			}
		}
	} else if hand1.handType > hand2.handType {
		result = 1
	} else if hand1.handType < hand2.handType {
		result = -1
	}

	return result
}

func sortHands(hands []hand) {
	slices.SortStableFunc(hands, compareHands)
}

func partOne(hands []hand) int {
	var totalWinnings int

	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}

	return totalWinnings
}

func compareHandsWithNewRules(hand1, hand2 hand) int {
	var result int

	cardPriority := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}

	if hand1.morphedHandType == hand2.morphedHandType {
		for i := 0; i < len(hand1.hand); i++ {
			if cardPriority[string(hand1.hand[i])] > cardPriority[string(hand2.hand[i])] {
				result = 1
				break
			} else if cardPriority[string(hand1.hand[i])] < cardPriority[string(hand2.hand[i])] {
				result = -1
				break
			}
		}
	} else if hand1.morphedHandType > hand2.morphedHandType {
		result = 1
	} else if hand1.morphedHandType < hand2.morphedHandType {
		result = -1
	}

	return result
}

func sortHandsWithNewRules(hands []hand) {
	slices.SortStableFunc(hands, compareHandsWithNewRules)
}

func partTwo(hands []hand) int {
	var totalWinnings int

	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}

	return totalWinnings
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := strings.Split(strings.TrimSpace(string(file)), "\n")
	processedHands := processHands(fileArray)

	// Part 1
	sortHands(processedHands)
	fmt.Printf("Part 1: %d\n", partOne(processedHands))

	// Part 2
	sortHandsWithNewRules(processedHands)
	fmt.Printf("Part 2: %d\n", partTwo(processedHands))
}
