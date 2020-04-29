package main

func numWays(steps int, arrLen int) int {
	mod := int(1e9) + 7
	dp := make([]int, min(steps+1, arrLen))
	dp[0] = 1
	for i := 0; i < steps; i++ {
		curr := make([]int, min(steps+1, arrLen))
		for j := 0; j < len(curr); j++ {
			curr[j] = dp[j]
			if j+1 < len(curr) {
				curr[j] += dp[j+1]
				curr[j] %= mod
			}
			if j-1 >= 0 {
				curr[j] += dp[j-1]
				curr[j] %= mod
			}
		}
		dp = curr
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
