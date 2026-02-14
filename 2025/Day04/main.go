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

	rows := [][]rune{}
	grid := strings.SplitSeq(strings.TrimSpace(string(file)), "\n")
	for row := range grid {
		rows = append(rows, []rune{})
		for _, roll := range row {
			rows[len(rows)-1] = append(rows[len(rows)-1], roll)
		}
	}

	return rows
}

func partOne(grid [][]rune) int {
	totalRolls := 0

	for i, row := range grid {
		for j, roll := range row {
			if roll == '@' {
				adjacentRolls := 0
				if i > 0 {
					for _, v := range grid[i-1][max(0, j-1):min(len(row), j+2)] {
						if v == '@' {
							adjacentRolls++
						}
					}
				}

				if i < len(grid)-1 {
					for _, v := range grid[i+1][max(0, j-1):min(len(row), j+2)] {
						if v == '@' {
							adjacentRolls++
						}
					}
				}

				if j > 0 && row[j-1] == '@' {
					adjacentRolls++
				}

				if j < len(grid[0])-1 && row[j+1] == '@' {
					adjacentRolls++
				}

				if adjacentRolls < 4 {
					totalRolls++
				}
			}
		}
	}

	return totalRolls
}

func partTwo(grid [][]rune) int {
	totalRolls := 0

	removedAny := true
	for removedAny {
		removedAny = false

		for i, row := range grid {
			for j, roll := range row {
				if roll == '@' {
					adjacentRolls := 0
					if i > 0 {
						for _, v := range grid[i-1][max(0, j-1):min(len(row), j+2)] {
							if v == '@' {
								adjacentRolls++
							}
						}
					}

					if i < len(grid)-1 {
						for _, v := range grid[i+1][max(0, j-1):min(len(row), j+2)] {
							if v == '@' {
								adjacentRolls++
							}
						}
					}

					if j > 0 && row[j-1] == '@' {
						adjacentRolls++
					}

					if j < len(grid[0])-1 && row[j+1] == '@' {
						adjacentRolls++
					}

					if adjacentRolls < 4 {
						grid[i][j] = '.'
						totalRolls++
						removedAny = removedAny || true
					} else {
						removedAny = removedAny || false
					}
				}
			}
		}
	}

	return totalRolls
}

func main() {
	var input [][]rune
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
