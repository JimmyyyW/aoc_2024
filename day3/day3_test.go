package day3

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	input := parseInput("input.txt")
	if len(input) < 1000 {
		t.Fail()
	}
}

func TestExamplePartOne(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := solve(input)
	if result != 161 {
		t.Fail()
	}
}

func TestExamplePartTwo(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	result := solvePartTwo(input)
	if result != 48 {
		t.Fail()
	}
}

func TestPartOne(t *testing.T) {
	input := parseInput("input.txt")
	result := solve(input)
	println(result)
}

func TestPartTwo(t *testing.T) {
	input := parseInput("input.txt")
	result := solvePartTwo(input)
	println(result)
}
