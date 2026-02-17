package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func timeIt(name string) func() {
	start := time.Now()
	return func() {
		fmt.Println(name, "took", time.Since(start))
	}
}

func parseInput() [][]rune {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "File open error: ", err)
		return nil
	}

	var input [][]rune
	for line := range strings.SplitSeq(strings.TrimSpace(string(file)), "\n") {
		input = append(input, []rune(line))
	}

	return input
}

type Loc struct {
	X int
	Y int
}

func partOne(input [][]rune) int {
	var split int

	var grid [][]rune
	grid = make([][]rune, 0, len(input))
	for _, v := range input {
		grid = append(grid, make([]rune, len(input[0])))
		copy(grid[len(grid)-1], v)
	}

	var start Loc
	for i, row := range grid {
		for j, c := range row {
			if c == 'S' {
				start = Loc{X: i + 1, Y: j}
			}
		}
	}

	var recurse func(loc Loc)
	recurse = func(loc Loc) {
		if loc.X == len(grid) || loc.Y == len(grid[0]) || loc.Y == -1 {
			return
		}

		if v := grid[loc.X][loc.Y]; v == '.' {
			grid[loc.X][loc.Y] = '|'
			recurse(Loc{X: loc.X + 1, Y: loc.Y})
		} else if v == '^' {
			recurse(Loc{X: loc.X, Y: loc.Y - 1})
			split++
			recurse(Loc{X: loc.X, Y: loc.Y + 1})
		}
	}

	recurse(start)

	return split
}

func partTwo(input [][]rune) int {
	var grid [][]rune
	grid = make([][]rune, 0, len(input))
	for _, v := range input {
		grid = append(grid, make([]rune, len(input[0])))
		copy(grid[len(grid)-1], v)
	}

	var start Loc
	for i, row := range grid {
		for j, c := range row {
			if c == 'S' {
				start = Loc{X: i + 1, Y: j}
			}
		}
	}

	cache := make(map[Loc]int)
	var recurse func(loc Loc) int
	recurse = func(loc Loc) int {
		if loc.X == len(grid) || loc.Y == len(grid[0]) || loc.Y == -1 {
			return 0
		}

		if val, found := cache[loc]; found {
			return val
		}

		var result int
		if v := grid[loc.X][loc.Y]; v == '.' {
			result = recurse(Loc{X: loc.X + 1, Y: loc.Y})
		} else {
			result += recurse(Loc{X: loc.X, Y: loc.Y - 1})
			result++
			result += recurse(Loc{X: loc.X, Y: loc.Y + 1})
		}

		cache[loc] = result
		return result
	}
	cache[start] = recurse(start) + 1

	return cache[start]
}

func main() {
	var input [][]rune
	func() {
		defer timeIt("Parse Input")()
		input = parseInput()
	}()

	func() {
		defer timeIt("Part One")()
		fmt.Println("Part One:", partOne(input))
	}()

	func() {
		defer timeIt("Part Two")()
		fmt.Println("Part Two:", partTwo(input))
	}()
}
