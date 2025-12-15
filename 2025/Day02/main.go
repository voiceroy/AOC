package main

import (
	"bytes"
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

func parseInput() [][2]int {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "File open error: ", err)
		return nil
	}

	var a, b int
	res := [][2]int{}
	lines := strings.SplitSeq(strings.TrimSpace(string(file)), ",")
	for numRange := range lines {
		_, _ = fmt.Sscanf(numRange, "%d-%d", &a, &b)
		res = append(res, [2]int{a, b})
	}

	return res
}

var powersOf10 = [...]int{
	1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000,
	1000000000, 10000000000, 100000000000, 1000000000000,
}

func partOne(ranges [][2]int) int {
	invalidIDs := 0
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			var digits int
			if i < 10 {
				digits = 1
			} else if i < 100 {
				digits = 2
			} else if i < 1000 {
				digits = 3
			} else if i < 10000 {
				digits = 4
			} else if i < 100000 {
				digits = 5
			} else if i < 1000000 {
				digits = 6
			} else if i < 10000000 {
				digits = 7
			} else if i < 100000000 {
				digits = 8
			} else {
				digits = 9
			}

			if digits%2 != 0 {
				continue
			}

			halfPower := powersOf10[digits/2]
			if i/halfPower == i%halfPower {
				invalidIDs += i
			}
		}
	}

	return invalidIDs
}

func partTwo(ranges [][2]int) int {
	invalidIDs := 0

	intBuffer := make([]byte, 0, 32)
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			intBuffer = intBuffer[:0]
			intBuffer = strconv.AppendInt(intBuffer, int64(i), 10)
			sLen := len(intBuffer)

			found := false
			for length := 1; length <= sLen/2; length++ {
				if sLen%length != 0 {
					continue
				}

				if bytes.Equal(intBuffer[length:], intBuffer[:sLen-length]) {
					found = true
					break
				}
			}

			if found {
				invalidIDs += i
			}
		}
	}

	return invalidIDs
}

func main() {
	var input [][2]int
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
