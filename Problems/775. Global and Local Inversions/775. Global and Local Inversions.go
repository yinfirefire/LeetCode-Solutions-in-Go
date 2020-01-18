package main

import "math"

func isIdealPermutation(A []int) bool {
	//all local inversions are global inversions
	//so if there are more global inversions, then from 0...i, there must be max(A[0...i])>A[i+2], because they cannot be neighbors
	max := A[0]
	for i := 0; i < len(A)-2; i++ {
		max = int(math.Max(float64(A[i]), float64(max)))
		if max > A[i+2] {
			return false
		}
	}
	return true
}
