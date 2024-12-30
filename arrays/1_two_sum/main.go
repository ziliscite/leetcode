package main

import "fmt"

var nums = []int{2, 7, 11, 15}
var target = 9

//
//Output: [0,1]

func main() {
	fmt.Println(twoSum(nums, target))
}

// Time complexity: O(n^2)
func twoSumBrute(nums []int, target int) []int {
	// iterating nums | O(n)
	for i := 0; i < len(nums); i++ {

		// iterating nums, while inside of the loop | O(n)
		for j := i + 1; j < len(nums); j++ {

			// if it sums up to the target
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

		// for each i, it takes O(n), meaning for n numbers, it takes O(n^2)
	}

	return nil
}

func twoSum(nums []int, target int) []int {
	// create maps of value and index | O(1)
	pairs := make(map[int]int)

	// []int{2, 7, 11, 15}
	// 9
	//
	// map[int]int{2: 0, 7: 1, 11: 2, 15: 3}

	// iterating nums | O(n)
	for i, v := range nums {
		// check if the target's pair (target - v) exists
		if j, ok := pairs[target-v]; ok {
			// returns the index
			return []int{j, i}
		}
		// if not, then add to the map
		pairs[v] = i
	}

	return nil
}
