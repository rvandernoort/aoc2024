package main

import (
	"fmt"
	"time"

	"rovervandernoort.nl/framework"
)

func createInitialChecksum(nums []int) *framework.Deque {
	index := 0
	checksum := framework.NewDeque()
	for i, num := range nums {
		for j := 0; j < num; j++ {
			if i%2 == 0 {
				checksum.PushBack(index)
			} else {
				checksum.PushBack(-1)
			}
		}
		if i%2 == 0 {
			index += 1
		}
	}
	return checksum
}

func reorderChecksum(checksum framework.Deque) []int {
	newChecksum := make([]int, checksum.Size())
	index := 0
	for checksum.Size() > 0 {
		front := checksum.PopFront().(int)
		if front == -1 {
			if checksum.Size() == 0 {
				break
			}
			back := checksum.PopBack().(int)
			for back == -1 {
				if checksum.Size() == 0 {
					break
				}
				back = checksum.PopBack().(int)
			}
			if back == -1 {
				break
			}
			newChecksum[index] = back
		} else {
			newChecksum[index] = front
		}
		index++
	}
	return newChecksum
}

func calculateChecksum(fileBlocks []int) (checksum int) {
	for i, num := range fileBlocks {
		checksum += num * i
	}
	return
}

func Part1(nums []int) (checksum int) {
	checksumList := createInitialChecksum(nums)
	reorderedChecksumList := reorderChecksum(*checksumList)
	checksum = calculateChecksum(reorderedChecksumList)
	return
}

// func Part1(blocks []int, empties []int, length int) (checksum int) {
// 	index := 0
// 	lastIndex := len(blocks)
// 	for i := 0; i < length; i++ {
// 		if i%2 == 0 {
// 			for j := 0; j < blocks[i/2]; j++ {
// 				checksum += index * i
// 				index += 1
// 			}
// 		} else {
// 			for j := 0; j < empties[(i-1)/2]; j++ {
// 				checksum += lastIndex * i
// 				index += 1
// 			}
// 		}
// 	}
// 	return
// }

// func Part1(blocks []int, empties []int, length int) int {
// 	checksum := 0
// 	block_nr := 0
// 	block_size := blocks[block_nr]
// 	block_start := 1
// 	block_type := "block"
// 	end_block_nr := len(blocks)-1
// 	end_block_size := blocks[end_block_nr]
// 	block_end := 1
// 	empty_nr := 0
// 	empty_size := empties[empty_nr]
// 	empty_start := 1
// 	for i := 0; i < length; i++ {
// 		if block_type == "block" {
// 			checksum += i * block_nr

// 			if block_start == block_size {
// 				block_nr++
// 				block_size = blocks[block_nr]
// 				block_start = 1
// 				block_type = "empty"
// 				continue
// 			}

// 			block_start++
// 		} else if block_type == "empty" {
// 			checksum += i * end_block_nr

// 			if empty_start == empty_size {
// 				empty_nr++
// 				empty_size = empties[empty_nr]
// 				empty_start = 1
// 				block_type = "block"
// 				continue
// 			}

// 			if block_end == end_block_size {
// 				end_block_nr--
// 				end_block_size = blocks[end_block_nr-1]
// 				block_end = 1
// 				block_type = "block"
// 				continue
// 			}

// 			block_end++
// 			empty_start++
// 		}
// 	}

// 	return checksum
// }

// func Part2(nums []uint64, target uint64) uint64 {

// 	return 0
// }

func main() {
	lines := framework.ReadData("input.txt")
	blocks := framework.ParseBlocks(lines[0])

	fmt.Printf("-----\nPart 1\n")
	start := time.Now()
	product := Part1(blocks)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

	// fmt.Printf("-----\nPart 2\n")
	// start = time.Now()
	// product = Part2(nums, target)
	// fmt.Printf("time: %v\n\n", time.Since(start))
	// fmt.Printf("output: %d\n", product)

}
