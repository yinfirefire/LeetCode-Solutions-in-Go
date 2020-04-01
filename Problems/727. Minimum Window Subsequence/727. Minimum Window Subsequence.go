package main

import "math"

func minWindow(S string, T string) string {
	m := len(S)
	n := len(T)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for j := range dp[0] {
		dp[0][j] = math.MaxInt32 / 2
	}
	for i := range dp {
		dp[i][0] = 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if S[i-1] == T[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}
	minLen := math.MaxInt32 / 2
	pos := -1
	// fmt.Println(dp)
	for i := 1; i <= m; i++ {
		if dp[i][n] < minLen {
			minLen = dp[i][n]
			pos = i
		}
	}
	if minLen >= math.MaxInt32/2 {
		return ""
	}
	return S[pos-minLen : pos]
}
