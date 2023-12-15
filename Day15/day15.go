package main

import (
	"fmt"
	"os"
	"regexp"
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

func findIndex(array []lens, label string) (int, bool) {
	for i := range array {
		if array[i].label == label {
			return i, true
		}
	}

	return -1, false
}

func partTwo(data string) int {
	var lensPower int
	var boxes [256][]lens

	pattern := regexp.MustCompile(`(\w+)([=-])(\d*)`)
	for _, seq := range strings.Split(data, ",") {
		match := pattern.FindStringSubmatch(seq)
		label := match[1]
		location := hash(label)

		switch match[2] {
		case "=":
			power, _ := strconv.Atoi(match[3])
			currentLens := lens{label, power}

			if index, found := findIndex(boxes[location], currentLens.label); found {
				boxes[location][index] = currentLens
			} else {
				boxes[location] = append(boxes[location], lens{label, power})
			}
		case "-":
			if index, found := findIndex(boxes[location], label); found {
				boxes[location] = append(boxes[location][:index], boxes[location][index+1:]...)
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
