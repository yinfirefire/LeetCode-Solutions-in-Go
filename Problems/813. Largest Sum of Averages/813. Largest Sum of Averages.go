package main

import "math"

func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)
	sums := make([]int, n)
	sums[0] = A[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + A[i]
	}
	dp := make([][]float64, K+1)
	for i := range dp {
		dp[i] = make([]float64, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[1][i] = float64(sums[i-1]) / float64(i)
	}
	for k := 2; k <= K; k++ {
		for i := k; i <= n; i++ {
			for j := i - 1; j >= k-1; j-- {
				dp[k][i] = math.Max(dp[k-1][j]+float64(sums[i-1]-sums[j-1])/float64(i-j), dp[k][i])
			}
		}
	}
	return dp[K][n]
}
