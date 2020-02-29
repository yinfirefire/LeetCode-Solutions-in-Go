package main

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, v := range strs {
		z := 0
		o := 0
		for _, c := range v {
			if c == '0' {
				z++
			} else {
				o++
			}
		}
		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {
				dp[i][j] = max(dp[i][j], dp[i-z][j-o]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
