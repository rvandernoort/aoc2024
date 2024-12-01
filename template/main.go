package main

import (
	"fmt"
	"time"

	"rovervandernoort.nl/framework"
)

func Part1(nums []uint64, target uint64) uint64 {

	return 0
}

func Part2(nums []uint64, target uint64) uint64 {

	return 0
}

func main() {
	lines := framework.ReadData("input.txt")
	nums := framework.ParseUints(lines)
	target := uint64(2020)

	fmt.Printf("-----\nPart 1\n")
	start := time.Now()
	product := Part1(nums, target)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

	fmt.Printf("-----\nPart 2\n")
	start = time.Now()
	product = Part2(nums, target)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

}
