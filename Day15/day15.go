package main

import (
	"fmt"
	"os"
	"strings"
)

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

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileString := strings.TrimSpace(string(file))

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(fileString))
}
