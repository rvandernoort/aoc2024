package main

import (
	"testing"

	"rovervandernoort.nl/framework"
)

func TestPart1(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := uint64(4)
	grid := framework.ParseGrid(input)
	got := Part1(grid)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart1Bigger(t *testing.T) {
	input := framework.ReadData("test_input_bigger.txt")
	want := uint64(18)
	grid := framework.ParseGrid(input)
	got := Part1(grid)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2Smaller(t *testing.T) {
	input := framework.ReadData("test_input_small2.txt")
	want := uint64(1)
	grid := framework.ParseGrid(input)
	got := Part2(grid)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := framework.ReadData("test_input_bigger.txt")
	want := uint64(9)
	grid := framework.ParseGrid(input)
	got := Part2(grid)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
