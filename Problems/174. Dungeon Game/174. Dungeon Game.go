package main

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	if dungeon[m-1][n-1] > 0 {
		dp[m-1][n-1] = 1
	} else {
		dp[m-1][n-1] = -dungeon[m-1][n-1] + 1
	}
	for j := n - 2; j >= 0; j-- {
		dp[m-1][j] = dp[m-1][j+1]
		if dungeon[m-1][j] > dp[m-1][j] {
			dp[m-1][j] = 1
		} else {
			dp[m-1][j] -= dungeon[m-1][j]
		}
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = math.MaxInt32
			if i < m-1 {
				dp[i][j] = min(dp[i][j], dp[i+1][j])
			}
			if j < n-1 {
				dp[i][j] = min(dp[i][j], dp[i][j+1])
			}
			if dungeon[i][j] > dp[i][j] {
				dp[i][j] = 1
			} else {
				dp[i][j] -= dungeon[i][j]
			}
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
