package main

func numOfArrays(n int, m int, K int) int {
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, K+1)
		}
	}
	mod := int(1e9) + 7
	//dp[i][j][k] -> the first ith elements, with j is the largest value, and current search cost is k
	//if arr[i] is the current largest, dp[i][j][k] += sum(dp[i-1][1....j-1][k-1]) -> arr[i]==j
	//if arr[i] is not the current largest, dp[i][j][k] += j*(dp[i-1][j][k] ->
	// arr[i]<=j and 1:i-1 elements have largest value j. There will be j kinds of choice for arr[i]
	for j := 1; j <= m; j++ {
		dp[1][j][1] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 1; k <= K; k++ {
				for l := 1; l < j-1; l++ {
					dp[i][j][k] += dp[i-1][l][k-1]
					dp[i][j][k] %= mod
				}
				dp[i][j][k] += j * dp[i-1][j][k]
				dp[i][j][k] %= mod
			}
		}
	}
	res := 0
	for j := 1; j < m; j++ {
		res += dp[n][j][K]
		res %= mod
	}
	return res
}
