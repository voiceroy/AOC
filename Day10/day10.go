package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	y int
	x int
}

func (current Point) Equal(other Point) bool {
	return current.x == other.x && current.y == other.y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func calcArea(points []Point) int {
	var totalArea int
	for i := 0; i < len(points)-1; i++ {
		totalArea += (points[i].x * points[i+1].y) - (points[i+1].x * points[i].y)
	}
	totalArea += (points[len(points)-1].x * points[0].y) - (points[0].x * points[len(points)-1].y)

	return abs(totalArea) / 2
}

func processGrid(data []string) Point {
	start := Point{0, 0}

outerLoop:
	for j := range data {
		for i := range data[j] {
			if string(data[j][i]) == "S" {
				start.y, start.x = j, i
				break outerLoop
			}
		}
	}

	pipes := []string{"|", "-", "L", "J", "7", "F"}
	// Can we can go up
	if start.y != 0 && slices.Contains([]string{"F", "|", "7"}, string(data[start.y-1][start.x])) {
		pipes = slices.DeleteFunc(pipes, func(str string) bool {
			return slices.Contains([]string{"7", "F", "-"}, str)
		})
	}

	// Can we can go down
	if start.x != len(data[start.y])-1 && slices.Contains([]string{"-", "J", "7"}, string(data[start.y][start.x+1])) {
		pipes = slices.DeleteFunc(pipes, func(str string) bool {
			return slices.Contains([]string{"|", "7", "J"}, str)
		})
	}

	// Can we can go left
	if start.x != 0 && slices.Contains([]string{"-", "L", "F"}, string(data[start.y][start.x-1])) {
		pipes = slices.DeleteFunc(pipes, func(str string) bool {
			return slices.Contains([]string{"|", "L", "F"}, str)
		})
	}

	// Can we go right
	if start.y != len(data)-1 && slices.Contains([]string{"|", "L", "J"}, string(data[start.y+1][start.x])) {
		pipes = slices.DeleteFunc(pipes, func(str string) bool {
			return slices.Contains([]string{"L", "J", "-"}, str)
		})
	}

	data[start.y] = strings.ReplaceAll(data[start.y], "S", pipes[0])
	return start
}

func partOne(data []string, start Point, seen *[]Point, directions map[string][][]int) int {
	var pathLength = 1

	previous, next := start, start
	next.y += directions[string(data[start.y][start.x])][0][0]
	next.x += directions[string(data[start.y][start.x])][0][1]
	for !start.Equal(next) {
		for _, direction := range directions[string(data[next.y][next.x])] {
			if !previous.Equal(Point{next.y + direction[0], next.x + direction[1]}) {
				previous = next
				next.y += direction[0]
				next.x += direction[1]
				*seen = append(*seen, previous)
				break
			}
		}

		pathLength++
	}

	return pathLength / 2
}

func partTwo(data []string, seen []Point) int {
	return calcArea(seen) + 1 - len(seen)/2
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := []string{}
	for _, row := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		fileArray = append(fileArray, row)
	}
	start := processGrid(fileArray)

	directions := map[string][][]int{
		"|": {{-1, 0}, {1, 0}},
		"-": {{0, 1}, {0, -1}},
		"L": {{-1, 0}, {0, 1}},
		"J": {{-1, 0}, {0, -1}},
		"7": {{0, -1}, {1, 0}},
		"F": {{0, 1}, {1, 0}},
	}
	seen := []Point{start}

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(fileArray, start, &seen, directions))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(fileArray, seen))
}
