package main

import (
	"fmt"
	"sort"

	"rovervandernoort.nl/framework"
)

func absDiffUint(x, y uint64) uint64 {
	if x < y {
		return y - x
	}
	return x - y
}

func Part1(left []uint64, right []uint64) uint64 {
	difference := uint64(0)

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	for i := 0; i < len(left); i++ {
		difference += absDiffUint(left[i], right[i])
	}

	return difference
}

func occurrences(nums []uint64, target uint64) uint64 {
	count := uint64(0)

	for _, num := range nums {
		if num == target {
			count++
		}
	}

	return count
}

func Part2(left, right []uint64) uint64 {
	similarity := uint64(0)

	for _, num := range left {
		occurrences := occurrences(right, num)
		similarity += num * occurrences
	}

	return similarity
}

func main() {
	lines := framework.ReadData("input.txt")
	left, right := framework.ParseLists(lines)

	fmt.Printf("-----\nPart 1\n")
	product := Part1(left, right)
	fmt.Printf("output: %d\n", product)

	fmt.Printf("-----\nPart 2\n")
	product = Part2(left, right)
	fmt.Printf("output: %d\n", product)
}
