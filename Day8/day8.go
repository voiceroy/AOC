package main

import (
	"fmt"
	"os"
	"strings"
)

func makeMap(data []string) map[string][]string {
	nodeMap := make(map[string][]string)

	currentNode, leftNode, rightNode := "", "", ""
	for _, mapping := range data {
		fmt.Sscanf(mapping, "%s = (%3s, %3s)", &currentNode, &leftNode, &rightNode)
		nodeMap[currentNode] = []string{leftNode, rightNode}
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

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func partTwo(directions string, nodeMap map[string][]string) int {
	var totalSteps = 1

	for node := range nodeMap {
		if strings.HasSuffix(node, "A") {
			currentNode := node

			for i := 0; ; i++ {
				if strings.HasSuffix(currentNode, "Z") {
					totalSteps = lcm(totalSteps, i)
					break
				}

				direction := directions[i%len(directions)]
				if direction == 76 {
					currentNode = nodeMap[currentNode][0]
				} else {
					currentNode = nodeMap[currentNode][1]
				}
			}
		}
	}

	return totalSteps
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

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(directions, nodeMap))
}
