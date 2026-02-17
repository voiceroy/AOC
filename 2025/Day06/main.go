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

type Homework struct {
	numbers    [][]string
	operations []rune
}

func parseInput() Homework {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "File open error: ", err)
		return Homework{}
	}

	input := Homework{}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	indexes := []int{}
	for i, v := range lines[len(lines)-1] {
		if v != ' ' {
			indexes = append(indexes, i)
			input.operations = append(input.operations, v)
		}
	}

	for _, line := range lines[:len(lines)-1] {
		input.numbers = append(input.numbers, []string{})

		for i := range indexes[:len(indexes)-1] {
			numLen := indexes[i+1] - indexes[i] - 1

			input.numbers[len(input.numbers)-1] = append(input.numbers[len(input.numbers)-1], line[indexes[i]:indexes[i]+numLen])
		}

		input.numbers[len(input.numbers)-1] = append(input.numbers[len(input.numbers)-1], line[indexes[len(indexes)-1]:])
	}

	return input
}

func partOne(input Homework) int {
	total := 0

	for i, op := range input.operations {
		var curTotal int
		switch op {
		case '*':
			{
				curTotal = 1
				for _, row := range input.numbers {
					num, _ := strconv.Atoi(strings.TrimSpace(row[i]))
					curTotal *= num
				}
			}
		case '+':
			{
				for _, row := range input.numbers {
					num, _ := strconv.Atoi(strings.TrimSpace(row[i]))
					curTotal += num
				}
			}
		}

		total += curTotal
	}

	return total
}

func partTwo(input Homework) int {
	total := 0

	for i, op := range input.operations {
		var curTotal, maxLen int

		col := make([]string, 0, len(input.numbers))
		for _, row := range input.numbers {
			col = append(col, row[i])
			maxLen = max(maxLen, len(row[i]))
		}

		nums := []int{}
		for j := range maxLen {
			curNum := 0
			for _, num := range col {
				if num[j] == ' ' {
					continue
				} else {
					curNum = curNum*10 + int(num[j]-'0')
				}
			}

			nums = append(nums, curNum)
		}

		switch op {
		case '*':
			{
				curTotal = 1
				for _, v := range nums {
					curTotal *= v
				}
			}
		case '+':
			{
				for _, v := range nums {
					curTotal += v
				}
			}
		}

		total += curTotal
	}

	return total
}

func main() {
	var input Homework
	func() {
		defer timeIt("Parse Input")()
		input = parseInput()
	}()

	if len(input.numbers) == 0 && len(input.operations) == 0 {
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
