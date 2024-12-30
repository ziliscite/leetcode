package main

// Not all distinct values in the range 1 to n
var nums = []int{4, 3, 2, 7, 8, 2, 3, 1}

var ans = []int{5, 6}

func main() {
	disappeared := findDisappearedNumbersBrute(nums)
	for i, _ := range ans {
		if ans[i] != disappeared[i] {
			panic("Wrong solution")
		}
	}
}

func findDisappearedNumbersBrute(nums []int) []int {
	// create maps O(1)
	all := make(map[int]bool)

	// remove duplicates O(n)
	for _, v := range nums {
		all[v] = true
	}

	// create maps O(1)
	sol := make([]int, 0)

	// iterate through nums again O(n)
	for i, _ := range nums {
		n := i + 1 // since it says range from 1 - n
		if _, ok := all[n]; !ok {
			// if not present, then it's the missing number
			sol = append(sol, n)
		}
	}

	return sol
}
