package main

import (
	"container/heap"
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

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func partOne(lists [][]int) int {
	totalDistance := 0

	heapOne, heapTwo := &IntHeap{}, &IntHeap{}

	*heapOne = append(*heapOne, lists[0]...)
	*heapTwo = append(*heapTwo, lists[1]...)

	heap.Init(heapOne)
	heap.Init(heapTwo)

	for range len(lists[0]) {
		totalDistance += abs(heap.Pop(heapOne).(int) - heap.Pop(heapTwo).(int))
	}

	return totalDistance
}

func partTwo(lists [][]int) int {
	similarityScore := 0

	count := make(map[int]int, len(lists[1]))
	for _, v := range lists[1] {
		count[v]++
	}

	for _, v := range lists[0] {
		similarityScore += v * count[v]
	}

	return similarityScore
}

func parseInput() [][]int {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "File open error: ", err)
		return nil
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	res := [][]int{{}, {}}
	for _, line := range lines {
		v := strings.Fields(line)

		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])

		res[0] = append(res[0], x)
		res[1] = append(res[1], y)
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
