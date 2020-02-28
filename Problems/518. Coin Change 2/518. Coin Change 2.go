package main

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 1
	for i := 1; i <= len(coins); i++ {
		dp[i][0] = 1
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j-coins[i-1] >= 0 {
				dp[i][j] += dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[len(coins)][amount]
}
