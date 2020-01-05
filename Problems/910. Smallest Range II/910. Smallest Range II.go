package main

import "sort"

func smallestRangeII(A []int, K int) int {
	//sort first and then add 0 or 2K to each position
	sort.Ints(A)
	min := A[0]
	max := A[len(A)-1]
	dif := max - min
	for i := 0; i < len(A)-1; i++ {
		if A[0]+2*K > A[i+1] {
			min = A[i+1]
		} else {
			min = A[0] + 2*K
		}
		if A[i]+2*K > max {
			max = A[i] + 2*K
		}
		if max-min < dif {
			dif = max - min
		}
	}
	return dif
}
