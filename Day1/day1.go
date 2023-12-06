package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func partOne(data string) int {
	calibrationValue := 0
	for _, line := range strings.Split(data, "\n") {
		leftMost := -1
		rightMost := -1

		for _, char := range line {
			if value, err := strconv.Atoi(string(char)); err == nil {
				if leftMost == -1 {
					leftMost = value
					rightMost = leftMost
				} else {
					rightMost = value
				}
			}
		}
		if leftMost != -1 {
			calibrationValue += leftMost*10 + rightMost
		}
	}

	return calibrationValue
}

func partTwo(data string) int {
	replaceMap := map[*regexp.Regexp]string{
		regexp.MustCompile("(o)(ne)"):   " 1 ",
		regexp.MustCompile("(t)(wo)"):   " 2 ",
		regexp.MustCompile("(thr)(ee)"): " 3 ",
		regexp.MustCompile("(fo)(ur)"):  " 4 ",
		regexp.MustCompile("(fi)(ve)"):  " 5 ",
		regexp.MustCompile("(s)(ix)"):   " 6 ",
		regexp.MustCompile("(sev)(en)"): " 7 ",
		regexp.MustCompile("(eig)(ht)"): " 8 ",
		regexp.MustCompile("(ni)(ne)"):  " 9 ",
	}

	for k, v := range replaceMap {
		data = k.ReplaceAllString(data, "$1"+v+"$2")
	}

	return partOne(data)
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot open file 'input'\n")
		return
	}
	fileString := string(file)

	// Part 1
	fmt.Printf("Part 1: %d\n", partOne(fileString))

	// Part 2
	fmt.Printf("Part 2: %d\n", partTwo(fileString))
}
