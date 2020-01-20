package main

import "math"

//dp[i][j] = number of results with i dices and the last number is j
//two situation
//1. i==rollMax[j]+1, e.g. i = 4, rollMax[0]=3, so 1111 is not allowed, then 1111 should be removed from the current result => dp[3][0]-=1
//2. i>rollMax[j]+1, e.g. i = 4, rollMax[0]=2, then x111(except 1111) should be removed from the current result
//	which is dp[4][0] -= (dp[4-rollMax[0]]-1][0~5]-dp[4-rollMax[0]-1][0])

func dieSimulator(n int, rollMax []int) int {
	mod := int(math.Pow(10, 9)) + 7
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 7)
	}
	for i := range dp[1] {
		dp[1][i] = 1
	}
	dp[1][6] = 6 //use dp[i][6] as sum of all results of i indices
	for i := 2; i <= n; i++ {
		sum := 0
		for j := 0; j <= 5; j++ {
			dp[i][j] = dp[i-1][6]
			if i == rollMax[j]+1 {
				dp[i][j] -= 1
			} else if i > rollMax[j]+1 {
				dp[i][j] -= dp[i-rollMax[j]-1][6] - dp[i-rollMax[j]-1][j]
				dp[i][j] = (dp[i][j] + mod) % mod
			}
			sum += dp[i][j]
			sum %= mod
		}
		dp[i][6] = sum
	}
	return dp[n][6]
}
