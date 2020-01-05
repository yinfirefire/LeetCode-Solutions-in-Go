package main

import "sort"

func orderlyQueue(S string, K int) string {
	//when K=1, we can rotate the array
	//e.g. 1,2,3,4,5 -> 2,3,4,5,1 -> 3,4,5,1,2
	//bubble sort is to compare and swap two neighbor number like a sliding window, and sliding n times
	//when K>=2
	//we can achieve the step swap, e.g.
	//5 2 1 3 4 -> 2 1 3 4 5 -> 2 3 4 5 1 -> 2 4 5 1 3 -> 2 5 1 3 4 (rotate the [1:] part)
	//so when K>=2 we just sort the string
	// if K==1, we find the smallest in all rotated strings
	ss := []byte(S)
	if K >= 2 {
		sort.Slice(ss, func(a, b int) bool {
			return ss[a] < ss[b]
		})
		return string(ss)
	} else {
		res := make([]string, len(ss))
		for i := 0; i < len(ss); i++ {
			res[i] = S[i:] + S[0:i]
		}
		sort.Strings(res)
		return res[0]
	}
}
