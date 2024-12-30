package main

import (
	"fmt"
	"time"

	"rovervandernoort.nl/framework"
)

func createInitialChecksum(nums []int) *framework.Deque[int] {
	index := 0
	checksum := framework.NewDeque[int]()
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

func reorderChecksum(checksum framework.Deque[int]) []int {
	newChecksum := make([]int, checksum.Size())
	index := 0
	for checksum.Size() > 0 {
		front, _ := checksum.PopFront()
		if front == -1 {
			if checksum.Size() == 0 {
				break
			}
			back, _ := checksum.PopBack()
			for back == -1 {
				if checksum.Size() == 0 {
					break
				}
				back, _ = checksum.PopBack()
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

func createFileSystem(blocks []int) framework.BlockDeque {
	fs := framework.NewBlockDeque()
	index := 0
	num := 0
	for i, block := range blocks {
		fs.PushBack(framework.FsBlock{Size: block, StartIndex: index, Number: num})
		index += block
		if i%2 == 0 {
			num = -1
		} else {
			num = (i + 1) / 2
		}
	}
	return *fs
}

func reorderFileSystem(fs framework.BlockDeque) (newFs framework.BlockDeque) {
	for fs.Size() > 0 {
		back, err := fs.PopBack()
		if !err {
			break
		}
		if back.Number == -1 || newFs.FirstValue() < back.Number {
			newFs.PushFront(back)
			continue
		} else {
			gap, i, err := fs.SearchFirstGap(back.Size)
			if !err {
				newFs.PushFront(back)
				continue
			}
			if gap.Size > back.Size {
				fs.InsertAt(i, framework.FsBlock{Size: gap.Size - back.Size, StartIndex: gap.StartIndex + back.Size, Number: -1})
			}
			//check at index to see if . blocks exists and extend either direction
			//or paste a new . block in place of the number block
			//no we need to change datastructure

			gap.Number = back.Number
			gap.Size = back.Size
			if gap.StartIndex+back.Size == back.StartIndex {
				for fs.Size() > 0 {
					back, err = fs.PopBack()
					if !err {
						return fs
					}
					newFs.PushFront(back)
				}
				return newFs
			}
		}
	}
	return newFs
}

func calculateChecksumBlock(fs framework.BlockDeque) (checksum int) {
	current, err := fs.PopFront()
	for err {
		if current.Number != -1 {
			for i := 0; i < current.Size; i++ {
				checksum += current.Number * (current.StartIndex + i)
			}
		}
		current, err = fs.PopFront()
	}
	return
}

func Part2(blocks []int) (checksum int) {
	fs := createFileSystem(blocks)
	fs = reorderFileSystem(fs)
	checksum = calculateChecksumBlock(fs)
	return
}

func main() {
	lines := framework.ReadData("input.txt")
	blocks := framework.ParseBlocks(lines[0])

	fmt.Printf("-----\nPart 1\n")
	start := time.Now()
	product := Part1(blocks)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

	fmt.Printf("-----\nPart 2\n")
	start = time.Now()
	product = Part2(blocks)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

}
