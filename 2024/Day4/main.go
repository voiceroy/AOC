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

func partOne(text [][]byte) int {
	xmas, n := 0, len(text)

	for row := range text {
		for col := range text[row][:len(text[row])-3] {
			if bytes.Equal(text[row][col:col+4], []byte("XMAS")) {
				xmas++
			}

			if bytes.Equal(text[row][n-col-4:n-col], []byte("SAMX")) {
				xmas++
			}
		}
	}

	for row := range text[:len(text)-3] {
		for col := range text {
			if bytes.Equal([]byte{text[row][col], text[row+1][col], text[row+2][col], text[row+3][col]}, []byte("XMAS")) {
				xmas++
			}

			if bytes.Equal([]byte{text[n-row-4][col], text[n-row-3][col], text[n-row-2][col], text[n-row-1][col]}, []byte("SAMX")) {
				xmas++
			}
		}
	}

	for k := 0; k < 2*n-1; k++ {
		for row := 0; row < n; row++ {
			col := k - row

			if col >= 0 && col < n && row+3 < n && col+3 < n {
				if bytes.Equal([]byte{text[row][col], text[row+1][col+1], text[row+2][col+2], text[row+3][col+3]}, []byte("XMAS")) {
					xmas++
				}

				if bytes.Equal([]byte{text[row][col], text[row+1][col+1], text[row+2][col+2], text[row+3][col+3]}, []byte("SAMX")) {
					xmas++
				}
			}
		}
	}

	for k := 0; k < 2*n-1; k++ {
		for row := 0; row < n; row++ {
			col := k - (n - 1 - row)

			if col >= 0 && col < n && row+3 < n && col-3 >= 0 {
				if bytes.Equal([]byte{text[row][col], text[row+1][col-1], text[row+2][col-2], text[row+3][col-3]}, []byte("XMAS")) {
					xmas++
				}

				if bytes.Equal([]byte{text[row][col], text[row+1][col-1], text[row+2][col-2], text[row+3][col-3]}, []byte("SAMX")) {
					xmas++
				}
			}
		}
	}

	return xmas
}

func partTwo(text [][]byte) int {
	xmas, n := 0, len(text)

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			if text[i][j] == 'A' {
				if text[i-1][j-1] == 'M' && text[i-1][j+1] == 'S' && text[i+1][j-1] == 'M' && text[i+1][j+1] == 'S' {
					xmas++
				} else if text[i-1][j-1] == 'M' && text[i-1][j+1] == 'M' && text[i+1][j-1] == 'S' && text[i+1][j+1] == 'S' {
					xmas++
				} else if text[i-1][j-1] == 'S' && text[i-1][j+1] == 'M' && text[i+1][j-1] == 'S' && text[i+1][j+1] == 'M' {
					xmas++
				} else if text[i-1][j-1] == 'S' && text[i-1][j+1] == 'S' && text[i+1][j-1] == 'M' && text[i+1][j+1] == 'M' {
					xmas++
				}
			}
		}
	}

	return xmas
}

func parseInput() [][]byte {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return bytes.Split(bytes.TrimRight(file, " \n"), []byte("\n"))
}

func main() {
	var input [][]byte

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
