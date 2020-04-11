package main

import "math"

func superEggDropTLE(K int, N int) int {
	// FIXME this O(kn^2) solution cause TLE
	// dp[k][n] is the steps to find F, if we have k eggs and n floors
	// if we drop at ith floor
	// if it broke, dp[k][n] = 1 + dp[k-1][i-1], because we know the F is lower and we lost an egg
	// if it did not break, dp[k][n] = 1 + dp[k][n-k], because we know F>=i+1, and we still have k eggs
	// so dp[k][n] = min(1+max(dp[k][i+1], dp[k-1][i-1])
	dp := make([][]int, 1+K)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	// dp[1][x] = x, we need to check from the bottom for every floor
	// dp[i][0] = 0, no need to check
	for j := 1; j <= N; j++ {
		dp[1][j] = j
	}
	for i := 2; i <= K; i++ {
		for j := 1; j <= N; j++ {
			dp[i][j] = math.MaxInt32
			for k := 1; k <= j; k++ {
				dp[i][j] = min(dp[i][j], 1+max(dp[i-1][k-1], dp[i][j-k]))
			}
		}
	}
	return dp[K][N]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
	To improve the time efficiency we focus on k
	since we are looking at max(dp[i-1][k-1], dp[i][j-k])
	dp[i-1][k-1] is increasing with k increasing
	dp[i][j-k] is decreasing with k increasing
	so dp[i][j] is minimum when the two function meet

	since in each iteration j is increasing, so in each iteration the k we need increasing
	we can keep the k in last iteration and start from there in the next round
*/

func superEggDrop(K int, n int) int {
	dp := make([][]int, K+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[1][i] = i
	}
	for i := 2; i <= K; i++ {
		k := 1
		for j := 1; j <= n; j++ {
			for k <= j && dp[i-1][k-1] < dp[i][j-k] {
				k++
			}
			dp[i][j] = 1 + dp[i-1][k-1]
		}
	}
	return dp[K][n]
}
