package main

import (
	"fmt"
	"strconv"
	"time"

	"rovervandernoort.nl/framework"
)

func combine(nums [][]int, result int, part2 bool) int {
	output := 0
	for _, num := range nums {
		if len(num) == 1 && num[0] == result {
			return result
		} else if len(num) == 1 {
			continue
		}
		newNums := [][]int{}
		newNum := []int{num[0]+num[1]}
		if len(num) >= 3 {
			newNums = append(newNums, append(newNum, num[2:]...))
		} else {
			newNums = append(newNums, newNum)
		}
		newNum = []int{num[0]*num[1]}
		if len(num) >= 3 {
			newNums = append(newNums, append(newNum, num[2:]...))
		} else {
			newNums = append(newNums, newNum)
		}
		if part2 {
			concat := fmt.Sprintf("%d%d", num[0], num[1])
			concatInt, err := strconv.Atoi(concat)
			if err != nil {
				fmt.Printf("%s is not a number\n", concat)
				panic(err)
			}
			newNum = []int{concatInt}
			if len(num) >= 3 {
				newNums = append(newNums, append(newNum, num[2:]...))
			} else {
				newNums = append(newNums, newNum)
			}
		}

		output = combine(newNums, result, part2)
		if output == result {
			return result
		}
	}
	return output
}

func Part1(equations []framework.Equation) int {
	count := 0
	for _, equation := range equations {
		nums := [][]int{}
		count += combine(append(nums, equation.Numbers), equation.Result, false)
	}

	return count
}

func Part2(equations []framework.Equation) int {
	count := 0
	for _, equation := range equations {
		nums := [][]int{}
		count += combine(append(nums, equation.Numbers), equation.Result, true)
	}

	return count
}

func main() {
	lines := framework.ReadData("input.txt")
	equations := framework.ParseEquations(lines)

	fmt.Printf("-----\nPart 1\n")
	start := time.Now()
	product := Part1(equations)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

	fmt.Printf("-----\nPart 2\n")
	start = time.Now()
	product = Part2(equations)
	fmt.Printf("time: %v\n\n", time.Since(start))
	fmt.Printf("output: %d\n", product)

}
