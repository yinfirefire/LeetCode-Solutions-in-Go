package main

func strangePrinter(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {
			if i == j {
				dp[i][j] = 1
			} else if j == i+1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = 2
				}
			} else {
				dp[i][j] = dp[i][i] + dp[i+1][j]
				for k := i; k < j; k++ {
					dp[i][j] = min(dp[i][k]+dp[k+1][j], dp[i][j])
				}
				if s[i] == s[j] {
					dp[i][j] -= 1
				}
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
