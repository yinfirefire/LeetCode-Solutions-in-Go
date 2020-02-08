package main

import "math"

func minFallingPathSum(arr [][]int) int {
	m := len(arr)
	n := len(arr[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	min1, min2 := math.MaxInt32, math.MaxInt32

	for i := range dp[0] {
		dp[0][i] = arr[0][i]
		if dp[0][i] <= min1 {
			min2 = min1
			min1 = dp[0][i]
		} else {
			min2 = min(min2, dp[0][i])
		}
	}

	for i := 1; i < m; i++ {
		curMin1, curMin2 := math.MaxInt32, math.MaxInt32
		for j := 0; j < n; j++ {
			if dp[i-1][j] == min1 {
				dp[i][j] = min2 + arr[i][j]
			} else {
				dp[i][j] = min1 + arr[i][j]
			}
			if dp[i][j] <= curMin1 {
				curMin2 = curMin1
				curMin1 = dp[i][j]
			} else {
				curMin2 = min(dp[i][j], curMin2)
			}
		}
		min1 = curMin1
		min2 = curMin2
	}
	return min1
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
