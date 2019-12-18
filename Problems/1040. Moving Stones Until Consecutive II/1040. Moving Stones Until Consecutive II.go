package main

import (
	"math"
	"sort"
)

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	// Upper bound:
	// 1, move the first stone A[0] to A[n-1]-n+1
	// the total number between new left bound and A[1] is the steps to take
	// A[n-1]-n+1-A[1]+1
	// 2, or move the last stone A[n-1] to A[0]+n-1
	// the total number between new right bound and A[n-2] is the steps to take
	// A[n-2]-(A[0]+n-1)+1
	//
	// Lower bound:
	// Sliding window size of len(stones), find vacancy in the current window, and the number of vacancy is the number
	// of steps
	// Edge Case: A[j]-A[i]==n-2 && j-i+1==n-1
	// e.g. 1,2,3,4,10: we need to move 1 to 6, then 10 to 5
	n := len(stones)
	i := 0
	min := float64(n)
	max := 0
	if stones[n-1]-n+1-stones[1]+1 < stones[n-2]-stones[0]-n+2 {
		max = stones[n-2] - stones[0] - n + 2
	} else {
		max = stones[n-1] - stones[1] - n + 2
	}
	for j := 0; j < n; j++ {
		for stones[j]-stones[i] >= n {
			i++
		}
		if stones[j]-stones[i] == n-2 && j-i == n-2 {
			min = math.Min(min, float64(2))
		} else {
			min = math.Min(min, float64(n-(j-i+1)))
		}
	}
	return []int{int(min), max}
}
