package main

import "math"

func mctFromLeafValues(arr []int) int {
	n := len(arr)
	mx := make([][]int, n)
	for i := range mx {
		mx[i] = make([]int, n)
	}
	for i := range arr {
		cur := arr[i]
		for j := i; j < n; j++ {
			if arr[j] > cur {
				cur = arr[j]
			}
			mx[i][j] = cur
		}
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			if i+1 == j {
				dp[i][j] = arr[i] * arr[j]
			} else {
				dp[i][j] = math.MaxInt32
				for k := i; k < j; k++ {
					dp[i][j] = Min(dp[i][j], dp[i][k]+dp[k+1][j]+mx[i][k]*mx[k+1][j])
				}
			}
		}
	}
	return dp[0][n-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
