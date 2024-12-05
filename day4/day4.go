package day4

import (
	"bufio"
	adventofcode2024 "github.com/jimmyyyw/aoc_2024"
	"regexp"
)

var pattern *regexp.Regexp = regexp.MustCompile(`XMAS|SAMX`)

func solve(input []string) int {
	sum := 0

	rowCount := rows(input)
	sum += rowCount

	columnCount := columns(input)
	sum += columnCount

	sum += diagRight(input)
	sum += diagLeft(input)
	return sum
}

func solvePartTwo(input []string) int {
	var mat [][]byte
	for i := range input {
		mat = append(mat, []byte(input[i]))
	}
	n, m := len(mat), len(mat[0])

	total := 0
	for r := 1; r < n-1; r++ {
		for c := 1; c < m-1; c++ {
			if mat[r][c] != 'A' {
				continue
			}
			str := string([]byte{mat[r-1][c-1], 'A', mat[r+1][c+1]})
			if str != "MAS" && str != "SAM" {
				continue
			}
			str = string([]byte{mat[r+1][c-1], 'A', mat[r-1][c+1]})
			if str != "MAS" && str != "SAM" {
				continue
			}
			total++
		}
	}
	return total
}

func rows(input []string) int {
	sum := 0
	for _, row := range input {
		sum += countMatches(row)
	}
	return sum
}

func columns(input []string) int {
	sum := 0
	for j := 0; j < len(input[0]); j++ {
		var column string
		for i := 0; i < len(input); i++ {
			column += string(input[i][j])
		}
		sum += countMatches(column)
	}
	return sum
}

func diagRight(input []string) int {
	rowCount := len(input)
	cols := len(input[0])

	var diagonals []string

	for col := 0; col <= cols-4; col++ {
		var diagonal string
		for r, c := 0, col; r < rowCount && c < cols; r, c = r+1, c+1 {
			diagonal += string(input[r][c])
		}
		if len(diagonal) >= 4 {
			diagonals = append(diagonals, diagonal)
		}
	}

	for row := 1; row <= rowCount-4; row++ {
		var diagonal string
		for r, c := row, 0; r < rowCount && c < cols; r, c = r+1, c+1 {
			diagonal += string(input[r][c])
		}
		if len(diagonal) >= 4 {
			diagonals = append(diagonals, diagonal)
		}
	}
	sum := 0
	for i := range diagonals {
		sum += countMatches(diagonals[i])
	}
	return sum
}

func diagLeft(input []string) int {
	rowCount := len(input)
	cols := len(input[0])

	var diagonals []string

	// Collect diagonals starting from the first row
	for col := cols - 1; col >= 3; col-- { // Start points in the first row, from the right to allow diagonals of length 4
		var diagonal string
		for r, c := 0, col; r < rowCount && c >= 0; r, c = r+1, c-1 {
			diagonal += string(input[r][c])
		}
		if len(diagonal) >= 4 {
			diagonals = append(diagonals, diagonal)
		}
	}

	// Collect diagonals starting from the first column (excluding the top-right corner)
	for row := 1; row <= rowCount-4; row++ { // Start points in the first column
		var diagonal string
		for r, c := row, cols-1; r < rowCount && c >= 0; r, c = r+1, c-1 {
			diagonal += string(input[r][c])
		}
		if len(diagonal) >= 4 {
			diagonals = append(diagonals, diagonal)
		}
	}

	sum := 0
	for i := range diagonals {
		sum += countMatches(diagonals[i])
	}
	return sum
}

func countMatches(input string) int {
	matches := 0
	start := 0
	for start < len(input) {
		loc := pattern.FindStringIndex(input[start:])
		if loc == nil {
			break
		}
		matches++
		start += loc[0] + 1
	}
	return matches
}

func parseInput(filename string) []string {
	var strings []string
	adventofcode2024.ReadToScanner(filename, func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			line := scanner.Text()
			strings = append(strings, line)
		}
	})
	return strings
}
