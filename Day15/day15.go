package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type lens struct {
	label       string
	focalLength int
}

func hash(str string) int {
	var currentHash int
	for _, char := range str {
		currentHash += int(char)
		currentHash *= 17
		currentHash %= 256
	}

	return currentHash
}

func partOne(data string) int {
	var hashSum int
	for _, seq := range strings.Split(data, ",") {
		hashSum += hash(seq)
	}

	return hashSum
}

func partTwo(data string) int {
	var lensPower int
	var boxes [256][]lens

	pattern := regexp.MustCompile(`(\w+)([=-])(\d*)`)
	for _, seq := range strings.Split(data, ",") {
		match := pattern.FindStringSubmatch(seq)
		label, location := match[1], hash(match[1])
		focalLength, _ := strconv.Atoi(match[3])
		currentLens := lens{label, focalLength}

		switch match[2] {
		case "=":
			if index := slices.IndexFunc(boxes[location], func(otherLens lens) bool {
				return currentLens.label == otherLens.label
			}); index >= 0 {
				boxes[location][index] = currentLens
			} else {
				boxes[location] = append(boxes[location], currentLens)
			}
		case "-":
			if index := slices.IndexFunc(boxes[location], func(otherLens lens) bool {
				return currentLens.label == otherLens.label
			}); index >= 0 {
				boxes[location] = slices.Delete(boxes[location], index, index+1)
			}
		}
	}

	for i, box := range boxes {
		for j, currentLens := range box {
			lensPower += (i + 1) * (j + 1) * currentLens.focalLength
		}
	}

	return lensPower
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileString := strings.TrimSpace(string(file))

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(fileString))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(fileString))
}
