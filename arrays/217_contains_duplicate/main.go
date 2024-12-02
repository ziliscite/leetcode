package main

import (
	"fmt"
	"sort"
)

var nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

func main() {
	contains := containsDuplicateBrute(nums)
	fmt.Println(contains)
}

// Time complexity: O(n log(n))
func containsDuplicateBrute(nums []int) bool {
	// sorting takes O(n log(n)) at the minimum
	sort.Ints(nums)

	// iterating takes O(n)
	for i, v := range nums[:len(nums)-1] {
		if v == nums[i+1] {
			return true
		}
	}

	return false
}

// Time complexity: O(n)
func containsDuplicateHash(nums []int) bool {
	// creating a map takes O(1)
	maps := make(map[int]bool)

	// iterating takes O(n)
	for _, v := range nums {
		// if the value exists in the map, then it is a duplicate
		_, ok := maps[v]
		if ok {
			return true
		}

		maps[v] = true
	}

	return false
}
