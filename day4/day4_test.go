package day4

import "testing"

func TestParseInput(t *testing.T) {
	count := len(parseInput("input.txt"))
	if count != 140 {
		t.Fail()
	}
}

func TestExamplePartOne(t *testing.T) {
	exampleInput := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	result := solve(exampleInput)
	println(result)
	if result != 18 {
		t.Fail()
	}
}

func TestPartOne(t *testing.T) {
	input := parseInput("input.txt")
	total := solve(input)
	println(total)
}

func TestPartTwo(t *testing.T) {
	input := parseInput("input.txt")
	total := solvePartTwo(input)
	println(total)
}

func TestCountMatches(t *testing.T) {
	matches := countMatches("MAMXMASAMX")
	if matches != 2 {
		t.Fail()
	}
}
