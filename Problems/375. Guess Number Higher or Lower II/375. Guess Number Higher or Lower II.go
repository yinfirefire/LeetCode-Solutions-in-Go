package main

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for j := 1; j <= n; j++ {
		for i := j - 1; i >= 1; i-- {
			dp[i][j] = min(i+dp[i+1][j], j+dp[i][j-1])
			//choose the leftmost or rightmost number
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], max(dp[i][k-1], dp[k+1][j])+k)
				//choose k, and then the extra cost is at most dp[i][k-1] or dp[k+1][j], because we know k is higher or lower
			}

		}
	}
	return dp[1][n]
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
