package main

import (
	"math"
	"sort"
)

func numMovesStones(a int, b int, c int) []int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	if arr[2]-arr[0] == 2 {
		return []int{2, 2}
	}
	dist1 := arr[2] - arr[1]
	dist2 := arr[1] - arr[0]
	min := 1
	max := dist1 + dist2 - 2
	if math.Min(float64(dist1), float64(dist2)) <= float64(2) {
		//e.g. 2,3,6; 2,4,6
		min = 1
	} else {
		min = 2
	}
	return []int{min, max}
}
