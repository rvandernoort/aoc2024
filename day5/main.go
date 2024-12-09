package main

import (
	"fmt"
	"time"

	"rovervandernoort.nl/framework"
)

func Part1(ordering map[uint64][]uint64, updates [][]uint64) (output uint64) {
	for _, row := range updates {
		correct := checkOrdering(ordering, row)
		if correct {
			middle := len(row) / 2
			output += row[middle]
		}
	}
	return
}

func checkOrdering(ordering map[uint64][]uint64, row []uint64) (correct bool) {
	correct = true
	for i := 0; i < len(row); i++ {
		orders := ordering[row[i]]
		for _, order := range orders {
			for j := 0; j < i; j++ {
				if row[j] == order {
					correct = false
				}
			}
		}
	}
	return
}

func checkLine(row []uint64, output uint64, ordering map[uint64][]uint64) uint64 {
	for i := 0; i < len(row); i++ {
		item := row[i]
		orders := ordering[item]
		for _, order := range orders {
			for j := 0; j < i; j++ {
				if row[j] == order {
					newRow := make([]uint64, len(row))
					copy(newRow, row)
					safe := newRow[i]
					newRow = framework.Remove(newRow, i)
					newRow = append(newRow[:i-1], append([]uint64{safe}, newRow[i-1:]...)...)

					if checkOrdering(ordering, newRow) {
						middle := len(newRow) / 2
						output += newRow[middle]
						return output
					} else {
						output += checkLine(newRow, output, ordering)
						return output
					}
				}
			}
		}
	}
	return output
}

func Part2(ordering map[uint64][]uint64, updates [][]uint64) (output uint64) {
	for _, row := range updates {
		correct := checkOrdering(ordering, row)
		if !correct {
			output += checkLine(row, 0, ordering)
		}
	}
	return
}

func main() {
	lines := framework.ReadData("input.txt")
	ordering, updates := framework.ParsePages(lines)

	fmt.Printf("-----\nPart 1\n")
	start := time.Now()
	product := Part1(ordering, updates)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

	fmt.Printf("-----\nPart 2\n")
	start = time.Now()
	product = Part2(ordering, updates)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)
}
