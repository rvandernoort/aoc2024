package main

import (
	"fmt"

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
	product := Part1(nums, target)
	fmt.Printf("output: %d\n", product)

	fmt.Printf("-----\nPart 2\n")
	product = Part2(nums, target)
	fmt.Printf("output: %d\n", product)
}
