package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func timer(name string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func partOne(text []byte) int {
	total, i, j := 0, 0, 0

outer:
	for i < len(text) {
		if nextMul := bytes.Index(text[i:], []byte("mul(")); nextMul >= 0 {
			i += nextMul + 3
		} else {
			break
		}

		numOne, numTwo := 0, 0

		j = i + 1
		for text[j] != ',' {
			if digit := text[j] - '0'; digit <= 9 {
				numOne = int(digit) + numOne*10
			} else {
				i = j + 1
				continue outer
			}

			j++
		}

		j++

		for text[j] != ')' {
			if digit := text[j] - '0'; digit <= 9 {
				numTwo = int(digit) + numTwo*10
			} else {
				i = j + 1
				continue outer
			}

			j++
		}

		total += numOne * numTwo
		i = j + 1
	}

	return total
}

func partTwo(text []byte) int {
	total, i, j := 0, 0, 0

	nextDont, nextMul := bytes.Index(text, []byte("don't()")), -1
outer:
	for i < len(text) {
		if nextMul = bytes.Index(text[i:], []byte("mul(")); nextMul >= 0 {
			i += nextMul + 3
		} else {
			break
		}

		if i > nextDont && nextDont != -1 {
			if nextDo := bytes.Index(text[i:], []byte("do()")); nextDo >= 0 {
				i += nextDo + 4
				nextDont = bytes.Index(text[i:], []byte("don't()")) + i
				continue
			} else {
				break
			}
		}

		j = i + 1
		numOne, numTwo := 0, 0
		for text[j] != ',' {
			if digit := text[j] - '0'; digit <= 9 {
				numOne = int(digit) + numOne*10
			} else {
				i = j + 1
				continue outer
			}

			j++
		}

		j++

		for text[j] != ')' {
			if digit := text[j] - '0'; digit <= 9 {
				numTwo = int(digit) + numTwo*10
			} else {
				i = j + 1
				continue outer
			}

			j++
		}

		total += numOne * numTwo
		i = j + 1
	}

	return total
}

func parseInput() []byte {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return bytes.TrimRight(file, " \n")
}

func main() {
	var input []byte

	func() {
		defer timer("Parsing")()
		input = parseInput()
	}()

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
