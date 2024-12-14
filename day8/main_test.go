package main

import (
	"testing"

	"rovervandernoort.nl/framework"
)

func TestPart1(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := int(14)
	antennas, xMax, yMax := framework.ParseAntennas(input)
	got := Part1(antennas, xMax, yMax)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart1_1(t *testing.T) {
	input := framework.ReadData("test_input_simple1.txt")
	want := int(2)
	antennas, xMax, yMax := framework.ParseAntennas(input)
	got := Part1(antennas, xMax, yMax)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart1_2(t *testing.T) {
	input := framework.ReadData("test_input_simple2.txt")
	want := int(4)
	antennas, xMax, yMax := framework.ParseAntennas(input)
	got := Part1(antennas, xMax, yMax)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart1_3(t *testing.T) {
	input := framework.ReadData("test_input_simple3.txt")
	want := int(4)
	antennas, xMax, yMax := framework.ParseAntennas(input)
	got := Part1(antennas, xMax, yMax)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := int(34)
	antennas, xMax, yMax := framework.ParseAntennas(input)
	got := Part2(antennas, xMax, yMax)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2_1(t *testing.T) {
	input := framework.ReadData("test_input_part2.txt")
	want := int(9)
	antennas, xMax, yMax := framework.ParseAntennas(input)
	got := Part2(antennas, xMax, yMax)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
