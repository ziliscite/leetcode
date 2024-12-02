package main

import (
	"fmt"
	"sort"
)

// All distinct values between 0 and n
var nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

func main() {
	//missing := missingNumberBrute(nums)

	//missing := []int{0, 1}
	fmt.Println(missingNumberSum(nums))
}

func missingNumberBrute(nums []int) int {
	// sorting takes O(n log(n))
	sort.Ints(nums)

	length := len(nums)

	// iterating takes O(n)
	for i := 0; i <= length; i++ {
		// if the value does not match the index, then it is a missing number
		if nums[i] != i {
			return i
		}

		// if there is no next number, then it is the missing number
		if nums[i] == length-1 {
			return i + 1
		}
	}

	return 0
}

func missingNumberSum(nums []int) int {
	length := len(nums)

	// the expected sum all the numbers between 0 and n
	sum := 0

	// summing takes O(n)
	for i := 0; i <= length; i++ {
		sum += i
	}

	actualSum := 0

	// another summing takes O(n)
	for _, v := range nums {
		actualSum += v
	}

	return sum - actualSum
}
