package day3

import (
	"bufio"
	adventofcode2024 "github.com/jimmyyyw/aoc_2024"
	"regexp"
	"strconv"
)

func solve(input string) int {
	regex := regexp.MustCompile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
	matches := regex.FindAllStringSubmatch(input, -1)
	sum := 0
	for i := range matches {
		println(matches[i][1])
		left, err := strconv.Atoi(matches[i][1])
		if err != nil {
			panic("left bad")
		}
		right, err := strconv.Atoi(matches[i][2])
		if err != nil {
			panic("right bad")
		}
		sum += left * right
	}
	return sum
}

func solvePartTwo(input string) int {
	regex := regexp.MustCompile("(mul\\(([0-9]{1,3}),([0-9]{1,3})\\)|don't\\(\\)|do\\(\\))")
	matches := regex.FindAllStringSubmatch(input, -1)
	sum := 0
	ignoreFlag := false
	for i := 0; i < len(matches); i++ {
		if matches[i][0] == "don't()" {
			ignoreFlag = true
			continue
		}

		if matches[i][0] == "do()" {
			ignoreFlag = false
			continue
		}
		if !ignoreFlag {
			left, err := strconv.Atoi(matches[i][2])
			if err != nil {
				panic("left bad")
			}
			right, err := strconv.Atoi(matches[i][3])
			if err != nil {
				panic("right bad")
			}
			sum += left * right
		}

	}
	return sum
}

func parseInput(filename string) string {
	var returnee string
	adventofcode2024.ReadToScanner(filename, func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			returnee += scanner.Text()
		}
	})
	return returnee
}
