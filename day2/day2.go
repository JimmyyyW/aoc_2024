package day2

import (
	"bufio"
	adventofcode2024 "github.com/jimmyyyw/aoc_2024"
	"strconv"
	"strings"
)

func solve(input [][]int) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += checkReport(input[i])
	}
	return sum
}

func checkReport(report []int) int {
	current := report[0]
	if current > report[1] {
		for i := 1; i < len(report); i++ {
			if !(current-report[i] > 0 && current-report[i] <= 3) {
				break
			} else {
				if i == len(report)-1 {
					return 1
				}
			}
			current = report[i]
		}
	}

	if current < report[1] {
		for i := 1; i < len(report); i++ {
			if !(current-report[i] < 0 && current-report[i] >= -3) {
				break
			} else {
				if i == len(report)-1 {
					return 1
				}
			}
			current = report[i]
		}
	}
	return 0
}

func solvePartTwo(input [][]int) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		if checkReport(input[i]) == 0 {
			for j := 0; j < len(input[i]); j++ {
				copySlice := make([]int, len(input[i]))
				copy(copySlice, input[i])
				removed := removeElement(copySlice, j)
				if checkReport(removed) == 1 {
					sum += 1
					break
				}
			}
		} else {
			sum += 1
		}
	}
	return sum
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice // Return the original slice if the index is out of bounds
	}
	return append(slice[:index], slice[index+1:]...)
}
func parseInput(filename string) [][]int {
	var reports [][]int
	adventofcode2024.ReadToScanner(filename, func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			var report []int
			line := scanner.Text()
			parts := strings.Fields(line)

			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err != nil {
					panic("grr")
				}
				report = append(report, num)
			}
			reports = append(reports, report)
		}
	})
	return reports
}
