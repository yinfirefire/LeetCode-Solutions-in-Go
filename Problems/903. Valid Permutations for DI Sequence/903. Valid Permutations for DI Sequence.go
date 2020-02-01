package main

func numPermsDISequence(S string) int {
	n := len(S)
	mod := int(1e9) + 7
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if S[i-1] == 'I' {
				for k := 0; k < j; k++ {
					dp[i][j] += dp[i-1][k]
					dp[i][j] %= mod
				}
			} else {
				for k := j; k <= i-1; k++ {
					dp[i][j] += dp[i-1][k]
					dp[i][j] %= mod
				}
			}
		}
	}
	res := 0
	for i := range dp[n] {
		res += dp[n][i]
		res %= mod
	}
	return res
}
