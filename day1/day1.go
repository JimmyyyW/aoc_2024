package main

import (
	"bufio"
	adventofcode2024 "github.com/jimmyyyw/aoc_2024"
	"strconv"
	"strings"
)

func similarity(left []int, right []int) int {
	var leftToMultiplier []pair
	for _, l := range left {

		multiplier := 0
		for _, r := range right {
			if l == r {
				multiplier += 1
			}
		}
		leftToMultiplier = append(leftToMultiplier, pair{l, multiplier})

	}

	sum := 0
	for _, el := range leftToMultiplier {
		sum += el.left * el.right
	}
	return sum
}

func solve(left []int, right []int) int {
	var pairs []pair
	var l int
	var r int
	for len(left) > 0 {
		left, l = smallest(left)
		right, r = smallest(right)
		pairs = append(
			pairs,
			pair{left: l, right: r},
		)
	}

	sum := 0
	for _, pair := range pairs {
		sum += diff(pair.left, pair.right)
	}
	return sum
}

func diff(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func smallest(arr []int) ([]int, int) {
	smallest := arr[0]
	index := 0
	for i, element := range arr {
		if element < smallest {
			smallest = arr[i]
			index = i
		}
	}
	newArr := append(arr[:index], arr[index+1:]...)
	return newArr, smallest
}

func parseInput(filename string) ([]int, []int, error) {
	var left []int
	var right []int

	adventofcode2024.ReadToScanner(filename, func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			line := scanner.Text()

			parts := strings.Fields(line)

			if len(parts) != 2 {
				panic("f*** you lines")
			}

			leftStr, err := strconv.Atoi(parts[0])
			if err != nil {
				panic("bad leftStr")
			}

			rightStr, err := strconv.Atoi(parts[1])
			if err != nil {
				panic("bad leftStr")
			}
			left = append(left, leftStr)
			right = append(right, rightStr)
		}
	})
	return left, right, nil
}

type pair struct {
	left  int
	right int
}
