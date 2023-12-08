package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func makeMap(data []string) map[string][]string {
	nodeMap := make(map[string][]string)
	pattern := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

	for _, mapping := range data {
		matches := pattern.FindAllStringSubmatch(mapping, -1)
		nodeMap[matches[0][1]] = []string{matches[0][2], matches[0][3]}
	}

	return nodeMap
}

func partOne(directions string, nodeMap map[string][]string) int {
	currentNode := "AAA"
	for i := 0; ; i++ {
		if currentNode == "ZZZ" {
			return i
		}

		direction := directions[i%len(directions)]
		if direction == 76 {
			currentNode = nodeMap[currentNode][0]
		} else {
			currentNode = nodeMap[currentNode][1]
		}
	}
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := strings.Split(strings.TrimSpace(string(file)), "\n")
	directions := fileArray[0]
	nodeMap := makeMap(fileArray[2:])

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(directions, nodeMap))
}
