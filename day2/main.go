package main

import (
	"fmt"
	"time"

	"rovervandernoort.nl/framework"
)

func checkSafety(line []uint64) (result bool) {
	result = true
	increasing := line[0] < line[1]
	for i := 0; i + 1 < len(line); i++ {
		diff := framework.AbsDiffUint(line[i], line[i+1])
		if diff < 1 || diff > 3 || (increasing && line[i] > line[i+1]) || (!increasing && line[i] < line[i+1]) {
			result = false
			break
		}
	}
	return
}

func Part1(nums [][]uint64) (result uint64) {
	for _, line := range nums {
		safe := checkSafety(line)
		if safe {
			result += 1
		}
	}
	return
}

func Part2(nums [][]uint64) (result uint64) {
	for _, line := range nums {
		safe := false
		safe = safe || checkSafety(line)
		if !safe {
			for i := 0; i < len(line); i++ {
				lineCopy := make([]uint64, len(line))
				copy(lineCopy, line)
				tmpLine := framework.Remove(lineCopy, i)
				safe = safe || checkSafety(tmpLine)
				if safe {
					break
				}
			}
		}
		if safe {
			result += 1
		}
	}
	return
}

func main() {
	lines := framework.ReadData("input.txt")
	nums := framework.ParseReport(lines)

	fmt.Printf("-----\nPart 1\n")
	start := time.Now()
	product := Part1(nums)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

	fmt.Printf("-----\nPart 2\n")
	start = time.Now()
	product = Part2(nums)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

}
