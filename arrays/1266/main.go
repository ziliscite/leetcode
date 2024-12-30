package main

import (
	"fmt"
	"math"
)

var points = [][]int{{1, 1}, {3, 4}, {-1, 0}}
var ans = 7

func main() {
	fmt.Println(minTimeToVisitAllPoints(points))
}

func minTimeToVisitAllPoints(points [][]int) int {
	sol := 0
	point := points[0] // pick initial point | O(1)

	// iterating through points | O(n)
	for _, v := range points[1:] {
		// the minimum number of moves is the maximum differences in x and y, because we can move diagonally
		time := math.Max(math.Abs(float64(point[0]-v[0])), math.Abs(float64(point[1]-v[1])))

		// copying point | O(1)
		copy(point, v)
		sol += int(time)
	}

	return sol
}
