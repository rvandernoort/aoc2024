package main

import (
	"fmt"
	"time"

	"rovervandernoort.nl/framework"
)

func Part1(grid [][]string) (count uint64) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]) - 3; j++ {
			if grid[i][j] == "X" && grid[i][j+1] == "M" && grid[i][j+2] == "A" && grid[i][j+3] == "S" {
				count += 1
			} else if grid[i][j] == "S" && grid[i][j+1] == "A" && grid[i][j+2] == "M" && grid[i][j+3] == "X" {
				count += 1
			}
		}
	}

	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid) - 3; i++ {
			if grid[i][j] == "X" && grid[i+1][j] == "M" && grid[i+2][j] == "A" && grid[i+3][j] == "S" {
				count += 1
			} else if grid[i][j] == "S" && grid[i+1][j] == "A" && grid[i+2][j] == "M" && grid[i+3][j] == "X" {
				count += 1
			}
		}
	}

	for i := 0; i < len(grid) - 3; i++ {
		for j := 0; j < len(grid[i]) - 3; j++ {
			if grid[i][j] == "X" && grid[i+1][j+1] == "M" && grid[i+2][j+2] == "A" && grid[i+3][j+3] == "S" {
				count += 1
			} else if grid[i][j] == "S" && grid[i+1][j+1] == "A" && grid[i+2][j+2] == "M" && grid[i+3][j+3] == "X" {
				count += 1
			}
		}
	}

	for i := 3; i < len(grid); i++ {
		for j := 0; j < len(grid[i]) - 3; j++ {
			if grid[i][j] == "X" && grid[i-1][j+1] == "M" && grid[i-2][j+2] == "A" && grid[i-3][j+3] == "S" {
				count += 1
			} else if grid[i][j] == "S" && grid[i-1][j+1] == "A" && grid[i-2][j+2] == "M" && grid[i-3][j+3] == "X" {
				count += 1
			}
		}
	}

	return
}
func Part2(grid [][]string) (count uint64) {
	for i := 0; i < len(grid) - 2; i++ {
		for j := 0; j < len(grid[i]) - 2; j++ {
			if grid[i+1][j+1] == "A" {
				if (grid[i][j] == "M" && grid[i+2][j+2] == "S" || grid[i][j] == "S" && grid[i+2][j+2] == "M") && (grid[i][j+2] == "M" && grid[i+2][j] == "S" || grid[i][j+2] == "S" && grid[i+2][j] == "M") {
					count += 1
				}
			}
		}
	}

	return
}

func main() {
	lines := framework.ReadData("input.txt")
	grid := framework.ParseGrid(lines)

	fmt.Printf("-----\nPart 1\n")
	start := time.Now()
	product := Part1(grid)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

	fmt.Printf("-----\nPart 2\n")
	start = time.Now()
	product = Part2(grid)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)
}
