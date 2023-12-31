package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne(data []string, scale int) int {
	var totalDistance, rowPsa, colPsa, galaxies = 0, []int{0}, []int{0}, [][]int{}
	for _, row := range data {
		if strings.Contains(row, "#") {
			rowPsa = append(rowPsa, 1+rowPsa[len(rowPsa)-1])
		} else {
			rowPsa = append(rowPsa, scale+rowPsa[len(rowPsa)-1])
		}
	}

	for i := 0; i < len(data[0]); i++ {
		colString := ""
		for _, row := range data {
			colString += string(row[i])
		}

		if strings.Contains(colString, "#") {
			colPsa = append(colPsa, 1+colPsa[len(colPsa)-1])
		} else {
			colPsa = append(colPsa, scale+colPsa[len(colPsa)-1])
		}
	}

	for i, row := range data {
		for j, col := range row {
			if string(col) == "#" {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

	for i, galaxy := range galaxies {
		for _, other := range galaxies[:i] {
			totalDistance += rowPsa[max(galaxy[0], other[0])] - rowPsa[min(galaxy[0], other[0])]
			totalDistance += colPsa[max(galaxy[1], other[1])] - colPsa[min(galaxy[1], other[1])]
		}
	}

	return totalDistance
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
