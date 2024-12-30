package main

import (
	"sort"
)

var nums = []int{8, 1, 2, 2, 3}

var output = []int{4, 0, 1, 1, 3}

func main() {
	smallerNumbersThanCurrent(nums)
}

func smallerNumbersThanCurrent(nums []int) []int {
	// copying the whole array | O(n)
	sorted := make([]int, len(nums))

	copy(sorted, nums)

	// sorting | O(n log(n))
	sort.Ints(sorted)

	// create maps of value and index | O(1)
	pairs := make(map[int]int)

	// iterating through sorted nums and adding to the map | O(n)
	for i, v := range sorted {
		// if its a duplicate, then we skip because
		// the index (numbers of smaller numbers) will be different if we didn't
		if _, ok := pairs[v]; ok {
			continue
		}
		pairs[v] = i
	}

	// iterating through nums | O(n)
	sol := make([]int, 0)
	for _, v := range nums {
		// append the index of the value as the number of smaller numbers
		sol = append(sol, pairs[v])
	}

	return sol
}
