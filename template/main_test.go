package main

import (
	"testing"

	"rovervandernoort.nl/framework"
)

func TestPart1(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := uint64(514579)
	nums := framework.ParseUints(input)
	got := Part1(nums, 2020)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := uint64(241861950)
	nums := framework.ParseUints(input)
	got := Part2(nums, 2020)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
