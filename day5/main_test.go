package main

import (
	"testing"

	"rovervandernoort.nl/framework"
)

func TestPart1(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := uint64(143)
	ordering, updates := framework.ParsePages(input)
	got := Part1(ordering, updates)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := uint64(241861950)
	ordering, updates := framework.ParsePages(input)
	got := Part2(ordering, updates)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
