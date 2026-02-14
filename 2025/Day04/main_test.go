package main

import "testing"

var input = parseInput()

func BenchmarkParsing(b *testing.B) {
	for b.Loop() {
		parseInput()
	}
}

func BenchmarkPartOne(b *testing.B) {
	for b.Loop() {
		partOne(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for b.Loop() {
		partTwo(input)
	}
}
