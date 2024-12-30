package main

import (
	"testing"

	"rovervandernoort.nl/framework"
)

func TestPart1(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := int(1928)
	blocks := framework.ParseBlocks(input[0])
	got := Part1(blocks)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart1_actual(t *testing.T) {
	input := framework.ReadData("input.txt")
	want := int(6446899523367)
	blocks := framework.ParseBlocks(input[0])
	got := Part1(blocks)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart1_S(t *testing.T) {
	input := framework.ReadData("test_input_small.txt")
	want := int(60)
	blocks := framework.ParseBlocks(input[0])
	got := Part1(blocks)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart1_B(t *testing.T) {
	input := framework.ReadData("test_input_bigger.txt")
	want := int(20)
	blocks := framework.ParseBlocks(input[0])
	got := Part1(blocks)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := framework.ReadData("test_input.txt")
	want := int(2858)
	blocks := framework.ParseBlocks(input[0])
	got := Part2(blocks)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2_S(t *testing.T) {
	input := framework.ReadData("test_input_blocks.txt")
	want := int(64)
	blocks := framework.ParseBlocks(input[0])
	got := Part2(blocks)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
