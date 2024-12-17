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

func ParsePages(input []string) (ordering map[uint64][]uint64, updates [][]uint64) {
	ordering = make(map[uint64][]uint64)
	part2 := false
	for _, line := range input {
		if line == "" {
			part2 = true
			continue
		}
		if !part2 {
			rule := strings.Split(line, "|")
			left, err := strconv.ParseUint(rule[0], 10, 64)
			if err != nil {
				fmt.Printf("%s is not a number\n", rule[0])
				panic(err)
			}
			right, err := strconv.ParseUint(rule[1], 10, 64)
			if err != nil {
				fmt.Printf("%s is not a number\n", rule[1])
				panic(err)
			}
			ordering[left] = append(ordering[left], right)
		}
		if part2 {
			lineList := []uint64{}
			for _, item := range strings.Split(line, ",") {
				num, err := strconv.ParseUint(string(item), 10, 64)
				if err != nil {
					fmt.Printf("%s is not a number\n", line)
					panic(err)
				}
				lineList = append(lineList, num)
			}
			updates = append(updates, lineList)
		}
	}
	return
}

type Tuple struct {
	X int
	Y int
}

func ParseAntennas(input []string) (map[string][]Tuple, int, int) {
	antennas := make(map[string][]Tuple)
	for i, line := range input {
		for j, letter := range line {
			strLetter := string(letter)
			if strLetter == "." || strLetter == "#" {
				continue
			}
			antennas[strLetter] = append(antennas[strLetter], Tuple{X: i, Y: j})
		}
	}
	return antennas, len(input[0]), len(input)
}

type Equation struct {
	Result int
	Numbers []int
}

func ParseEquations(input []string) []Equation {
	equations := make([]Equation, 0)
	for _, line := range input {
		equation := strings.Split(line, ": ")
		result, err := strconv.Atoi(equation[0])
		if err != nil {
			fmt.Printf("%s is not a number\n", equation[0])
			panic(err)
		}
		numbers := strings.Split(equation[1], " ")
		nums := []int{}
		for _, num := range numbers {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("%s is not a number\n", num)
				panic(err)
			}
			nums = append(nums, numInt)
		}
		equations = append(equations, Equation{Result: result, Numbers: nums})
	}
	return equations
}

