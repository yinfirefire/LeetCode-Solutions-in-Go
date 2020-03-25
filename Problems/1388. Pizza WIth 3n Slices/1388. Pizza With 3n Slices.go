package main

func maxSizeSlices(slices []int) int {
	return max(helper(slices[0:len(slices)-1]), helper(slices[1:len(slices)]))
}

func helper(slices []int) int {
	m := len(slices)
	n := (m + 1) / 3
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[1][1] = slices[0]
	for i := 2; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-2][j-1]+slices[i-1])
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// same as house robber ii
