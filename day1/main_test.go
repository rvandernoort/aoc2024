package main

import (
	"testing"

	"rovervandernoort.nl/framework"
)

func TestPart1(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := uint64(11)
	left, right := framework.ParseLists(input)
	got := Part1(left, right)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := uint64(31)
	left, right := framework.ParseLists(input)
	got := Part2(left, right)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
