package main

import (
	"fmt"
	"time"

	"rovervandernoort.nl/framework"
)

func Part1(ordering map[uint64][]uint64, updates [][]uint64) (output uint64) {
	for _, row := range updates {
		correct := true
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

func Part2(ordering map[uint64][]uint64, updates [][]uint64) (output uint64) {
	for _, row := range updates {
		correct := true
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
		if !correct {
			for i := 0; i < len(row); i++ {
				orders := ordering[row[i]]
				for _, order := range orders {
					for j := 0; j < i; j++ {
						if row[j] == order {
							newRow := make([]uint64, len(row))
							copy(newRow, row)
							newRow = framework.Remove(newRow, i)
							newRow = append(newRow[:j+1], append([]uint64{row[i]}, newRow[j+1:]...)...)
							if checkOrdering(ordering, newRow) {
								middle := len(newRow) / 2
								output += newRow[middle]
								return
							}
						}
					}
				}
			}
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
