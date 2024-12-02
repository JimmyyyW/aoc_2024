package day2

import (
	"testing"
)

func TestExample(t *testing.T) {
	var input = [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := solve(input)
	if result != 2 {
		println("result should be 2 but was")
		println(result)
		t.Fail()
	}
}

func TestExamplePartTwo(t *testing.T) {
	var input = [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := solvePartTwo(input)
	if result != 4 {
		println("result should be 4 but was")
		println(result)
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

func TestInputParser(t *testing.T) {
	result := parseInput("input.txt")
	if len(result) != 1000 {
		t.Fail()
	}
}
