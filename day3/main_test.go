package main

import (
	"testing"

	"rovervandernoort.nl/framework"
)

func TestPart1(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := uint64(161)
	got := Part1(input)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart1double(t *testing.T) {
	input := framework.ReadData("test_input_double.txt")
	want := uint64(322)
	got := Part1(input)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := framework.ReadData("test_input_do_dont.txt")
	want := uint64(48)
	got := Part2(input)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2double(t *testing.T) {
	input := framework.ReadData("test_input_do_dont_double.txt")
	want := uint64(96)
	got := Part2(input)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}