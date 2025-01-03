package main

import (
	"fmt"
	"slices"
	"time"

	"rovervandernoort.nl/framework"
)

func Part1(left []uint64, right []uint64) uint64 {
	difference := uint64(0)

	slices.Sort(left)
	slices.Sort(right)

	for i := 0; i < len(left); i++ {
		difference += framework.AbsDiffUint(left[i], right[i])
	}

	return difference
}

func Occurences(nums []uint64, target uint64) uint64 {
	frequencyMap := make(map[uint64]uint64)
	for _, num := range nums {
		frequencyMap[num]++
	}
	return frequencyMap[target]
}

// func Occurrences(nums []uint64, target uint64) uint64 {
// 	count := uint64(0)

// 	for _, num := range nums {
// 		if num == target {
// 			count++
// 		}
// 	}

// 	return count
// }

func Part2(left, right []uint64) uint64 {
	similarity := uint64(0)

	for _, num := range left {
		occurrences := Occurences(right, num)
		similarity += num * occurrences
	}

	return similarity
}

func main() {
	lines := framework.ReadData("input.txt")
	left, right := framework.ParseLists(lines)

	fmt.Printf("-----\nPart 1\n")
	start := time.Now()
	product := Part1(left, right)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

	fmt.Printf("-----\nPart 2\n")
	start = time.Now()
	product = Part2(left, right)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

}
