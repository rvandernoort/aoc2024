package main

import (
	"testing"

	"rovervandernoort.nl/framework"
)

func TestPart1(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := 3749
	nums := framework.ParseEquations(input)
	got := Part1(nums)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := 11387
	nums := framework.ParseEquations(input)
	got := Part2(nums)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
