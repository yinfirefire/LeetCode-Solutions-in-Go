package main

func minCut(s string) int {
	n := len(s)
	dp := make([]int, n, n)
	is := make([][]bool, n)
	for i := range is {
		is[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		min := i
		for j := i; j >= 0; j-- {
			if s[i] == s[j] && (i-j < 2 || is[j+1][i-1]) {
				is[j][i] = true
				if j == 0 {
					min = 0
				} else {
					if min > dp[j-1]+1 {
						min = dp[j-1] + 1
					}
				}
			}
		}
		dp[i] = min
	}
	return dp[n-1]
}
