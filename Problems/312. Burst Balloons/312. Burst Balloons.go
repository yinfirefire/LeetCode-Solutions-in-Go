package main

func maxCoins(nums []int) int {
	n := len(nums)
	nn := make([]int, n+2)
	nn[0], nn[n+1] = 1, 1
	for i := 0; i < n; i++ {
		nn[i+1] = nums[i]
	}
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for j := 1; j <= n; j++ {
		for i := j; i >= 1; i-- {
			for k := i; k <= j; k++ {
				cur := nn[k] * nn[i-1] * nn[j+1]
				prev := dp[i][k-1] + dp[k+1][j]
				dp[i][j] = max(prev+cur, dp[i][j])
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
