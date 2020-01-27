package main

import "math"

func minDifficulty(jobDifficulty []int, d int) int {
	//let dp[i][j] is the result for given i day to finish the [1:j] job
	//we need to decide what is the difficulty at the ith day
	//if the ith day we finish job[k:j], then ith day difficulty is max(jobDifficulty[k:j])
	//then dp[i][j] = dp[i-1][k-1]+max(jobDifficulty[k:j])
	length := len(jobDifficulty)
	dp := make([][]int, d+1)
	for i := range dp {
		dp[i] = make([]int, length+1)
	}
	curMax := 0
	for i := 1; i <= length; i++ {
		curMax = max(jobDifficulty[i-1], curMax)
		dp[1][i] = curMax
	}
	for i := 2; i <= d; i++ {
		for j := 1; j <= length; j++ {
			if i > j {
				dp[i][j] = -1
				continue
			}
			dp[i][j] = math.MaxInt32
			nextDay := 0
			for k := j; k > 1; k-- {
				nextDay = max(jobDifficulty[k-1], nextDay)
				if dp[i-1][k-1] != -1 {
					dp[i][j] = min(dp[i][j], dp[i-1][k-1]+nextDay)
				}
			}
		}
	}
	return dp[d][length]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
