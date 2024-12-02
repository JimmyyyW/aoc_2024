package main

import "testing"

func TestParseInput(t *testing.T) {
	left, right, err := parseInput("input.txt")
	if err != nil {
		t.Fail()
	}
	if len(left) != 1000 {
		t.Fail()
	}
	if len(right) != 1000 {
		t.Fail()
	}
}

func TestExample(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	if solve(left, right) != 11 {
		t.Fail()
	}
}

func TestPartTwoExample(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	if similarity(left, right) != 31 {
		t.Fail()
	}
}

func TestSolution(t *testing.T) {
	left, right, err := parseInput("input.txt")
	if err != nil {
		t.Fail()
	}
	result := solve(left, right)
	println(result)
}

func TestPartTwo(t *testing.T) {
	left, right, err := parseInput("input.txt")
	if err != nil {
		t.Fail()
	}
	result := similarity(left, right)
	println(result)
}

func TestExampleMk2(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	if sortThenDiff(left, right) != 11 {
		t.Fail()
	}
}

func TestSolutionMk2(t *testing.T) {
	left, right, err := parseInput("input.txt")
	if err != nil {
		t.Fail()
	}
	resultTakeTwo := sortThenDiff(left, right)
	println(resultTakeTwo)
	if resultTakeTwo != 1579939 {
		t.Fail()
	}
}
