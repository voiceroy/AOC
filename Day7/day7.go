package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Hand struct {
	hand            string
	handType        int
	morphedHandType int
	bid             int
}

func morphHand(hand string) string {
	if !strings.Contains(hand, "J") {
		return hand
	}

	var frequencies = make(map[rune]int)
	for _, card := range hand {
		frequencies[card]++
	}

	var maxCard rune
	var maxCount int
	for card, count := range frequencies {
		if maxCount < count && card != 74 {
			maxCard = card
			maxCount = count
		}
	}

	if maxCard == 0 {
		return hand
	}
	return strings.ReplaceAll(hand, "J", string(maxCard))
}

func getHandRank(hand string) int {
	var handRank int

	var cardCount = make(map[rune]int)
	for _, card := range hand {
		cardCount[card]++
	}

	var counts []int
	for _, count := range cardCount {
		counts = append(counts, count)
	}

	slices.Sort(counts)
	if slices.Compare(counts, []int{5}) == 0 {
		handRank = 7
	} else if slices.Compare(counts, []int{1, 4}) == 0 {
		handRank = 6
	} else if slices.Compare(counts, []int{2, 3}) == 0 {
		handRank = 5
	} else if slices.Compare(counts, []int{1, 1, 3}) == 0 {
		handRank = 4
	} else if slices.Compare(counts, []int{1, 2, 2}) == 0 {
		handRank = 3
	} else if slices.Compare(counts, []int{1, 1, 1, 2}) == 0 {
		handRank = 2
	} else if slices.Compare(counts, []int{1, 1, 1, 1, 1}) == 0 {
		handRank = 1
	}

	return handRank
}

func processHands(data []string) []Hand {
	var processedHands []Hand

	var bid int
	var hand string
	for _, line := range data {
		fmt.Sscanf(line, "%s %d\n", &hand, &bid)
		morphedHand := morphHand(hand)
		processedHands = append(processedHands, Hand{hand, getHandRank(hand), getHandRank(morphedHand), bid})
	}

	return processedHands
}

func genCompareHandFunc(cardStrength string, newRule bool) func(a, b Hand) int {
	return func(a, b Hand) int {
		var aHandType, bHandType = a.handType, b.handType
		if newRule {
			aHandType, bHandType = a.morphedHandType, b.morphedHandType
		}

		if aHandType > bHandType {
			return 1
		} else if aHandType == bHandType {
			for i := 0; i < len(a.hand); i++ {
				switch cmp.Compare(strings.Index(cardStrength, string(a.hand[i])), strings.Index(cardStrength, string(b.hand[i]))) {
				case 1:
					return 1
				case -1:
					return -1
				}
			}
		}

		return -1
	}
}

func partOne(hands []Hand) int {
	var totalWinnings int

	compareFunc := genCompareHandFunc("23456789TJQKA", false)
	slices.SortStableFunc(hands, compareFunc)
	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
	}

	return totalWinnings
}

func partTwo(hands []Hand) int {
	var totalWinnings int

	compareFunc := genCompareHandFunc("J23456789TQKA", true)
	slices.SortStableFunc(hands, compareFunc)
	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
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
	fmt.Printf("Part 1: %d\n", partOne(processedHands))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(processedHands))
}
