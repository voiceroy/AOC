package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func timer(name string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func remove(slice []int, i int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:i]...)
	return append(newSlice, slice[i+1:]...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}

	return 1
}

func partOne(lists [][]int) int {
	safe := 0
outer:
	for _, list := range lists {
		if increasing := list[1] - list[0]; increasing != 0 && abs(increasing) >= 1 && abs(increasing) <= 3 {
			increasing = sign(increasing)
			for i := range list[1 : len(list)-1] {
				if curDiff := list[i+2] - list[i+1]; curDiff != 0 && sign(curDiff) == increasing {
					if !(abs(curDiff) >= 1 && abs(curDiff) <= 3) {
						continue outer
					}
				} else {
					continue outer
				}

			}

			safe++
		}
	}

	return safe
}

func partTwo(lists [][]int) int {
	safe := 0
	for _, list := range lists {
		if partOne([][]int{list}) == 1 {
			safe++
		} else {
			for i := range list {
				if partOne([][]int{remove(list, i)}) == 1 {
					safe++
					break
				}
			}
		}
	}

	return safe
}

func parseInput() [][]int {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "File open error: ", err)
		return nil
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	res := make([][]int, 0, len(lines))
	for _, line := range lines {
		nums := strings.Fields(line)

		list := make([]int, 0, len(nums))
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			list = append(list, n)
		}

		res = append(res, list)
	}

	return res
}

func main() {
	input := parseInput()
	if input == nil {
		return
	}

	// Part One
	func() {
		defer timer("Part One")()
		one := partOne(input)
		println("Part One:", one)
	}()

	// Part Two
	func() {
		defer timer("Part Two")()
		two := partTwo(input)
		println("Part Two:", two)
	}()
}
