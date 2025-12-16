package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func timeIt(name string) func() {
	start := time.Now()
	return func() {
		fmt.Println(name, "took", time.Since(start))
	}
}

func parseInput() [][]int {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "File open error: ", err)
		return nil
	}

	banks := [][]int{}
	lines := strings.SplitSeq(strings.TrimSpace(string(file)), "\n")
	for bank := range lines {
		banks = append(banks, []int{})
		for _, battery := range bank {
			banks[len(banks)-1] = append(banks[len(banks)-1], int(battery-'0'))
		}
	}

	return banks
}

func getMaxIndex(s []int, idx int) int {
	i, v := 0, s[0]
	for j, u := range s {
		if v < u {
			v = u
			i = j
		}
	}

	return i + idx
}

func partOne(banks [][]int) int {
	totalJoltage := 0
	for _, bank := range banks {
		x := getMaxIndex(bank[:len(bank)-1], 0)
		y := slices.Max(bank[x+1:])
		totalJoltage += bank[x]*10 + y
	}

	return totalJoltage
}

func partTwo(banks [][]int) int {
	totalJoltage := 0

	stack := make([]int, 0, len(banks[0]))
	for _, bank := range banks {
		stack = stack[:0]
		toRemove := len(bank) - 12
		for _, battery := range bank {
			for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < battery {
				stack = stack[:len(stack)-1]
				toRemove--
			}

			stack = append(stack, battery)
		}

		currentJoltage := 0
		for i := range 12 {
			currentJoltage = currentJoltage*10 + stack[i]
		}

		totalJoltage += currentJoltage
	}

	return totalJoltage
}

func main() {
	var input [][]int
	func() {
		defer timeIt("Parse Input")()
		input = parseInput()
	}()

	if input == nil {
		return
	}

	func() {
		defer timeIt("Part One")()
		fmt.Println("Part One:", partOne(input))
	}()

	func() {
		defer timeIt("Part Two")()
		fmt.Println("Part Two:", partTwo(input))
	}()
}
