package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func timeIt(name string) func() {
	start := time.Now()
	return func() {
		fmt.Println(name, "took", time.Since(start))
	}
}

type Input struct {
	ranges      [][2]int
	ingredients []int
}

func parseInput() Input {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "File open error: ", err)
		return Input{}
	}

	input := Input{}
	lines := strings.SplitSeq(string(file), "\n")
	for line := range lines {
		if len(line) == 0 {
			break
		} else {
			parts := strings.Split(line, "-")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			input.ranges = append(input.ranges, [2]int{x, y})
		}
	}

	for line := range lines {
		val, _ := strconv.Atoi(line)
		input.ingredients = append(input.ingredients, val)
	}

	return input
}

func partOne(input Input) int {
	fresh := 0

	for _, ingredient := range input.ingredients {
		for _, rng := range input.ranges {
			if rng[0] <= ingredient && ingredient <= rng[1] {
				fresh++
				break
			}
		}
	}

	return fresh
}

func partTwo(input Input) int {
	slices.SortFunc(input.ranges, func(i, j [2]int) int {
		return i[0] - j[0]
	})

	merged := make([][2]int, 0, len(input.ranges))
	currentRange := input.ranges[0]
	for _, v := range input.ranges[1:] {
		if v[0] < currentRange[1] {
			if v[1] > currentRange[1] {
				currentRange[1] = v[1]
			}
		} else {
			merged = append(merged, currentRange)
			currentRange = v
		}
	}
	merged = append(merged, currentRange)

	totalFresh := 0
	for _, v := range merged {
		totalFresh += v[1] - v[0] + 1
	}

	return totalFresh
}

func main() {
	var input Input
	func() {
		defer timeIt("Parse Input")()
		input = parseInput()
	}()

	if len(input.ingredients) == 0 && len(input.ranges) == 0 {
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
