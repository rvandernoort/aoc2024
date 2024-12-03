package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"rovervandernoort.nl/framework"
)

func Part1(input []string) uint64 {
	muls := uint64(0)
	for _, line := range input {
		working := true
		for working {
			mulIndex := strings.Index(line, "mul(")
			if mulIndex == -1 {
				working = false
				break
			}
			start := mulIndex + 4
			end := strings.IndexByte(line[start:], ')') + start
			content := line[start:end]
			nums := strings.Split(content, ",")
			if len(nums) != 2 {
				line = line[start+1:]
				continue
			}
			a, err := strconv.ParseUint(nums[0], 10, 64)
			if err != nil {
				line = line[start+1:]
				continue
			}
			b, err := strconv.ParseUint(nums[1], 10, 64)
			if err != nil {
				line = line[start+1:]
				continue
			}
			muls += a * b
			line = line[end+1:]
		}
	}
	return muls
}

func Part2(input []string) uint64 {
	muls := uint64(0)
  enabled := true
	for _, line := range input {
		working := true
		for working {
			mulIndex := strings.Index(line, "mul(")
			doIndex := strings.Index(line, "do()")
			dontIndex := strings.Index(line, "don't()")
			if mulIndex == -1 {
				working = false
				break
			}
			if doIndex != dontIndex {
				if doIndex == -1 {
					if dontIndex < mulIndex {
						enabled = false
						line = line[dontIndex+7:]
						continue
					}
				} else if dontIndex == -1 {
					if doIndex < mulIndex {
						enabled = true
						line = line[doIndex+4:]
						continue
					}
				} else if doIndex < mulIndex && doIndex < dontIndex {
					enabled = true
					line = line[doIndex+4:]
					continue
				} else if dontIndex < mulIndex && dontIndex < doIndex {
					enabled = false
					line = line[dontIndex+7:]
					continue
				}
			}
			start := mulIndex + 4
			end := strings.IndexByte(line[start:], ')') + start
			if enabled {
				content := line[start:end]
				nums := strings.Split(content, ",")
				if len(nums) != 2 {
					line = line[start+1:]
					continue
				}
				a, err := strconv.ParseUint(nums[0], 10, 64)
				if err != nil {
					line = line[start+1:]
					continue
				}
				b, err := strconv.ParseUint(nums[1], 10, 64)
				if err != nil {
					line = line[start+1:]
					continue
				}
				muls += a * b
			}
			line = line[end+1:]
		}
	}
	return muls
}

func main() {
	lines := framework.ReadData("input.txt")

	fmt.Printf("-----\nPart 1\n")
	start := time.Now()
	nums := Part1(lines)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", nums)


	lines = framework.ReadData("input.txt")
	fmt.Printf("-----\nPart 2\n")
	start = time.Now()
	nums = Part2(lines)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", nums)
}
