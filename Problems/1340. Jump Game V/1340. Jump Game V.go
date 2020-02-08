package main

import "sort"

func maxJumps(arr []int, d int) int {
	res := 1
	l := len(arr)
	pairs := make([][]int, len(arr))
	for i := range arr {
		pairs[i] = []int{arr[i], i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	dp := make([]int, len(arr))
	for i := range pairs {
		idx := pairs[i][1]
		cur := 1
		for j := 1; j <= d; j++ {
			if idx+j < l {
				if arr[idx+j] >= arr[idx] {
					break
				}
				cur = max(cur, dp[idx+j]+1)
			}
		}
		for j := 1; j <= d; j++ {
			if idx-j >= 0 {
				if arr[idx-j] >= arr[idx] {
					break
				}
				cur = max(cur, dp[idx-j]+1)
			}
		}
		dp[idx] = cur
		res = max(dp[idx], res)
	}
	// fmt.Println(dp)
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
