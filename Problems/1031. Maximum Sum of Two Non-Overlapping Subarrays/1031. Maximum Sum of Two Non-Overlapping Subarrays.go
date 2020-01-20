package main

func maxSumTwoNoOverlap(A []int, L, M int) int {
	prefix := make([]int, len(A)+1)
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + A[i-1]
	}
	return max(find(prefix, L, M), find(prefix, M, L))
}

func find(prefix []int, L, M int) int {
	maxL := 0
	maxTotal := 0
	for i := L + M; i < len(prefix); i++ {
		//fix M length subarray at the last M elements of current subarray
		//L might be just on the left of M, or may be several step more to the left
		maxL = max(maxL, prefix[i-M]-prefix[i-M-L])
		maxTotal = max(maxTotal, maxL+prefix[i]-prefix[i-M])
	}
	return maxTotal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
