package main

import "sort"

func sortThenDiff(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i, el := range left {
		diff := el - right[i]
		if right[i] > el {
			diff = right[i] - el
		}
		sum += diff
	}
	return sum
}
