package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInts(values []string) []int {
	var result []int
	for _, v := range values {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}

	return result
}

func parseIntsFromLines(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		result = append(result, parseInts(strings.Fields(line)))
	}

	return result
}

func generateMaps(data []string) [][][]int {
	var mapArray = make([][][]int, 7)

	for _, section := range data {
		lines := strings.Split(section, "\n")
		switch {
		case strings.HasPrefix(lines[0], "seed-to-soil map"):
			mapArray[0] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "soil-to-fertilizer map"):
			mapArray[1] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "fertilizer-to-water map"):
			mapArray[2] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "water-to-light map"):
			mapArray[3] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "light-to-temperature map"):
			mapArray[4] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "temperature-to-humidity map"):
			mapArray[5] = parseIntsFromLines(lines[1:])

		case strings.HasPrefix(lines[0], "humidity-to-location map"):
			mapArray[6] = parseIntsFromLines(lines[1:])
		}
	}

	return mapArray
}

func getDestination(mapping *[][]int, src int) int {
	for _, row := range *mapping {
		if row[1] <= src && src < row[1]+row[2] {
			return src + row[0] - row[1]
		}
	}

	return src
}

func partOne(seeds *[]int, maps *[][][]int) int {
	lowestValue := math.MaxInt
	for _, value := range *seeds {
		for _, mapping := range *maps {
			value = getDestination(&mapping, value)
		}

		if value < lowestValue {
			lowestValue = value
		}
	}

	return lowestValue
}

func generateSeedRanges(seeds *[]int) [][]int {
	var seedRanges [][]int
	for i := 0; i < len(*seeds); i += 2 {
		seedRanges = append(seedRanges, []int{(*seeds)[i], (*seeds)[i] + (*seeds)[i+1]})
	}

	return seedRanges
}

func partTwo(seeds *[]int, maps *[][][]int) int {
	var seedRanges = generateSeedRanges(seeds)

	for _, mapping := range *maps {
		var splitRanges [][]int
		for len(seedRanges) > 0 {
			seedRange := seedRanges[len(seedRanges)-1]
			seedRanges = seedRanges[:len(seedRanges)-1]

			var broken bool
			seedStart, seedEnd := seedRange[0], seedRange[1]
			for _, interval := range mapping {
				broken = false
				overlapStart := max(seedStart, interval[1])
				overlapEnd := min(seedEnd, interval[1]+interval[2])
				if overlapStart < overlapEnd {
					splitRanges = append(splitRanges, []int{overlapStart - interval[1] + interval[0], overlapEnd - interval[1] + interval[0]})
					if overlapStart > seedStart {
						seedRanges = append(seedRanges, []int{seedStart, overlapStart})
					}
					if seedEnd > overlapEnd {
						seedRanges = append(seedRanges, []int{overlapEnd, seedEnd})
					}
					broken = true
					break
				}
			}

			if !broken {
				splitRanges = append(splitRanges, []int{seedStart, seedEnd})
			}
		}

		seedRanges = splitRanges
	}

	return slices.MinFunc(seedRanges, func(a, b []int) int { return cmp.Compare(a[0], b[0]) })[0]
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := strings.Split(strings.TrimSpace(string(file)), "\n\n")
	var seeds = parseInts(strings.Fields(strings.Split(fileArray[0], ":")[1]))
	var mappings [][][]int = generateMaps(fileArray)

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(&seeds, &mappings))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(&seeds, &mappings))
}
