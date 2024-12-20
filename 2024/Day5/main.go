package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func timer(name string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func partOne(input Input) int {
	total := 0

outer:
	for _, update := range input.updates {
		for i, page := range update {
			for _, other := range input.rules[page] {
				if slices.Contains(update[:i], other) {
					continue outer
				}
			}
		}

		total += update[len(update)/2]
	}

	return total
}

func partTwo(input Input) int {
	total := 0

outer:
	for _, update := range input.updates {
		for i, page := range update {
			for _, other := range input.rules[page] {
				if slices.Contains(update[:i], other) {
					slices.SortFunc(update, func(i, j int) int {
						if slices.Contains(input.rules[i], j) {
							return -1
						} else {
							return 1
						}
					})

					total += update[len(update)/2]
					continue outer
				}
			}
		}
	}

	return total
}

type Input struct {
	rules   map[int][]int
	updates [][]int
}

func parseInput() Input {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	input := Input{make(map[int][]int), [][]int{}}

	var a, b, k int
	for i, row := range lines {
		if row == "" {
			k = i
			break
		} else {
			fmt.Sscanf(row, "%d|%d", &a, &b)
			input.rules[a] = append(input.rules[a], b)
		}
	}

	for _, row := range lines[k+1:] {
		nums := strings.FieldsFunc(row, func(r rune) bool {
			if r == ',' {
				return true
			} else {
				return false
			}
		})

		array := []int{}
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			array = append(array, n)
		}

		input.updates = append(input.updates, array)
	}

	input.updates = input.updates[:len(input.updates)-1]
	return input
}

func main() {
	var input Input

	func() {
		defer timer("Parsing")()
		input = parseInput()
	}()

	// Part One
	func() {
		defer timer("Part One")()
		one := partOne(input)
		println("Part One:", one)
	}()

	// Part Two
	func() {
		defer timer("Part Two")()
		two := partTwo(input)
		println("Part Two:", two)
	}()
}
