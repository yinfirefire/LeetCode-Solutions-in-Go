package main

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = map[int]int{}
	}
	res := 0
	for i := range A {
		for j := 0; j < i; j++ {
			dist := A[i] - A[j]
			res += dp[j][dist]
			dp[i][dist] += dp[j][dist] + 1
		}
	}
	return res
}
