package framework

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func ParseUints(input []string) []uint64 {
	var result []uint64
	for _, line := range input {
		num, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			fmt.Printf("%s is not a number\n", line)
			panic(err)
		}
		result = append(result, num)
	}
	return result
}

func ParseReport(input []string) (result [][]uint64) {
	for _, line := range input {
		lineList := []uint64{}
		for _, item := range strings.Split(line, " ") {
			num, err := strconv.ParseUint(string(item), 10, 64)
			if err != nil {
				fmt.Printf("%s is not a number\n", line)
				panic(err)
			}
			lineList = append(lineList, num)
		}
		result = append(result, lineList)
	}
	return
}

func ParseLists(input []string) ([]uint64, []uint64) {
	var result []uint64
	var result2 []uint64
	for _, line := range input {
		nums := strings.Split(line, "   ")
		if len(nums) >= 2 {
			num1, err := strconv.ParseUint(nums[0], 10, 64)
			if err != nil {
				fmt.Printf("%s is not a number\n", nums[0])
				panic(err)
			}
			num2, err := strconv.ParseUint(nums[1], 10, 64)
			if err != nil {
				fmt.Printf("%s is not a number\n", nums[1])
				panic(err)
			}
			result = append(result, num1)
			result2 = append(result2, num2)
		}
	}
	return result, result2
}

func ParseGrid(input []string) (grid [][]string) {
	for _, line := range input {
		lineGrid := []string{}
		for _, letter := range line {
			strLetter := string(letter)
			lineGrid = append(lineGrid, strLetter)
		}
		grid = append(grid, lineGrid)
	}
	return
}
