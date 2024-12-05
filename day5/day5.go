package day5

import (
	"bufio"
	adventofcode2024 "github.com/jimmyyyw/aoc_2024"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`(\d{1,2})\|(\d{1,2})`)

func solve(pageMap map[int][]int, updates [][]int) (int, int) {
	sum := 0
	sumBad := 0
	for _, update := range updates {
		isValid := true
		for i, page := range update {

			if i == len(update)-1 {
				break
			}

			if !obeysPageOrder(page, update[i+1], pageMap) {
				isValid = false
				break
			}

		}
		if isValid {
			sum += middle(update)
		} else {
			sort.SliceStable(update, func(i, j int) bool {
				return obeysPageOrder(update[i], update[j], pageMap)
			})
			sumBad += middle(update)
		}

	}
	return sum, sumBad
}

func obeysPageOrder(a int, b int, rules map[int][]int) bool {
	return slices.Contains(rules[a], b)
}
func middle(updates []int) int {
	midIndex := len(updates) / 2
	return updates[midIndex]
}

func parseInput(filename string) (map[int][]int, [][]int) {
	var pageMap = make(map[int][]int)
	var updates = make([][]int, 0)
	adventofcode2024.ReadToScanner(filename, func(scanner *bufio.Scanner) {
		secondSection := false
		for scanner.Scan() {
			line := scanner.Text()
			if line == "\n" || line == "" {
				secondSection = true
				continue
			}
			if !secondSection {
				matches := pattern.FindStringSubmatch(line)
				left, errL := strconv.Atoi(matches[1])
				right, errR := strconv.Atoi(matches[2])
				if errL != nil || errR != nil {
					panic("shiz")
				}
				pageMap[left] = append(pageMap[left], right)
			}

			if secondSection {
				var nums []int
				strNums := strings.Split(line, ",")
				for i := range strNums {
					num, err := strconv.Atoi(strNums[i])
					if err != nil {
						panic("couldn't parse section two num")
					}
					nums = append(nums, num)
				}
				updates = append(updates, nums)
			}
		}
	})
	return pageMap, updates
}
