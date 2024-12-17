package main

import (
	"testing"

	"rovervandernoort.nl/framework"
)

func TestPart1(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := int(1928)
	blocks, empties, length := framework.ParseBlocks(input[0])
	got := Part1(blocks, empties, length)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart1_S(t *testing.T) {
	input := framework.ReadData("test_input_small.txt")
	want := int(60)
	blocks, empties, length := framework.ParseBlocks(input[0])
	got := Part1(blocks, empties, length)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// func TestPart2(t *testing.T) {
// 	input := framework.ReadData("test_input.txt")
// 	want := uint64(241861950)
// 	nums := framework.ParseUints(input)
// 	got := Part2(nums, 2020)
// 	if got != want {
// 		t.Errorf("got %d, want %d", got, want)
// 	}
// }
