package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func timeIt(name string) func() {
	start := time.Now()
	return func() {
		fmt.Println(name, "took", time.Since(start))
	}
}

func parseInput() []int {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "File open error: ", err)
		return nil
	}

	res := []int{}
	lines := strings.SplitSeq(strings.TrimSpace(string(file)), "\n")
	for line := range lines {
		val, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			res = append(res, -val)
		} else {
			res = append(res, val)
		}
	}

	return res
}

func partOne(rotations []int) int {
	zeroed := 0

	dial := 50
	for _, v := range rotations {
		dial = (dial + v) % 100
		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			zeroed++
		}
	}

	return zeroed
}

func mod(x, y int) int {
	if x < 0 {
		return (x + y) % y
	}

	return x % y
}

func partTwo(rotations []int) int {
	zeroed, dial := 0, 50
	for _, v := range rotations {
		if v < 0 {
			zeroed += -v / 100
			v = -(-v % 100)

			if dial != 0 && dial+v <= 0 {
				zeroed++
			}
		} else {
			zeroed += v / 100
			v %= 100

			if dial+v >= 100 {
				zeroed++
			}
		}

		dial = mod(dial+v, 100)
	}

	return zeroed
}

func main() {
	var input []int
	func() {
		defer timeIt("Parse Input")()
		input = parseInput()
	}()

	if input == nil {
		return
	}

	func() {
		defer timeIt("Part One")()
		fmt.Println("Part One:", partOne(input))
	}()

	func() {
		defer timeIt("Part Two")()
		fmt.Println("Part Two:", partTwo(input))
	}()
}
