package main

func findDerangement(n int) int {
	dp := make([]int, n+1)
	if n < 2 {
		return 0
	}
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	mod := int(1e9) + 7
	for i := 3; i <= n; i++ {
		// for the ith number, there are i-1 number already, the ith number has i-1 choice of placement
		// for where to place the ith number, there are two situations
		// 1. ith number exchange with jth number, jth number go to ith index, ith number go to jth index
		// now the other i-2 numbers has dp[i-2] derangement
		// 2. ith number go to jth index, but jth number did not go to ith index,
		// so now there are i-1 numbers, and for each number there is an invalid position (for the jth number, the ith index is invalid)
		// so now there will be dp[i-1] derangement
		dp[i] = (i - 1) * (dp[i-1] + dp[i-2]) % mod
	}
	return dp[n]
}
