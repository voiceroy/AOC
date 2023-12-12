package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type location struct {
	x int
	y int
}

func partOne(data []string, scale int) int {
	var sumLengths int
	var emptyRows, emptyColumns []int
	var galaxyLocations []location

	for i, row := range data {
		if !strings.Contains(row, "#") {
			emptyRows = append(emptyRows, i)
		}
	}

outerLoop:
	for j := range data[0] {
		for i := range data {
			if string(data[i][j]) == "#" {
				continue outerLoop
			}
		}

		emptyColumns = append(emptyColumns, j)
	}

	for i, row := range data {
		for j, char := range row {
			if string(char) == "#" {
				galaxyLocations = append(galaxyLocations, location{i, j})
			}
		}
	}

	for i, loc := range galaxyLocations {
		for _, loc2 := range galaxyLocations[:i] {
			for r := min(loc.x, loc2.x); r < max(loc.x, loc2.x); r++ {
				if slices.Contains(emptyRows, r) {
					sumLengths += scale
				} else {
					sumLengths++
				}
			}

			for c := min(loc.y, loc2.y); c < max(loc.y, loc2.y); c++ {
				if slices.Contains(emptyColumns, c) {
					sumLengths += scale
				} else {
					sumLengths++
				}
			}
		}
	}

	return sumLengths
}

func partTwo(data []string, scale int) int {
	return partOne(data, scale)
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := strings.Split(strings.TrimSpace(string(file)), "\n")

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(fileArray, 2))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(fileArray, 1000000))
}
