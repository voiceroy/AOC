package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int64
	y int64
}

func getArea(vertices []Point) int64 {
	var totalArea int64
	for i := 0; i < len(vertices)-1; i++ {
		totalArea += (vertices[i].x * vertices[i+1].y) - (vertices[i+1].x * vertices[i].y)
	}
	totalArea += (vertices[len(vertices)-1].x * vertices[0].y) - (vertices[0].x * vertices[len(vertices)-1].y)

	return int64(math.Abs(float64(totalArea)) / 2)
}

func partOne(data []string) int64 {
	var coordinates []Point
	var pathLength int64

	var steps int64
	var direction string
	var currentCoordinate = Point{0, 0}
	for _, line := range data {
		_, _ = fmt.Sscanf(line, "%s %d", &direction, &steps)

		switch direction {
		case "U":
			currentCoordinate.y += steps
		case "R":
			currentCoordinate.x += steps
		case "L":
			currentCoordinate.x -= steps
		case "D":
			currentCoordinate.y -= steps
		}

		pathLength += steps
		coordinates = append(coordinates, currentCoordinate)
	}

	return (getArea(coordinates) - pathLength/2 + 1) + pathLength
}

func partTwo(data []string) int64 {
	var coordinates []Point
	var pathLength int64

	var currentCoordinate = Point{0, 0}
	for _, line := range data {
		hexCode := strings.Fields(line)[2]
		hexCode = hexCode[2 : len(hexCode)-1]
		step, _ := strconv.ParseInt(hexCode[:5], 16, 64)

		switch string(hexCode[len(hexCode)-1]) {
		case "0":
			currentCoordinate.x += step
		case "1":
			currentCoordinate.y -= step
		case "2":
			currentCoordinate.x -= step
		case "3":
			currentCoordinate.y += step
		}

		pathLength += step
		coordinates = append(coordinates, currentCoordinate)
	}

	return (getArea(coordinates) - pathLength/2 + 1) + pathLength
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open 'input'")
		return
	}

	fileArray := strings.Split(strings.TrimSpace(string(file)), "\n")

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(fileArray))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(fileArray))
}
