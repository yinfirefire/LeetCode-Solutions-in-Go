package main

func countVowelPermutation(n int) int {
	ht := make(map[int][]int)
	ht[0] = []int{0, 1, 0, 0, 0}
	ht[1] = []int{1, 0, 1, 0, 0}
	ht[2] = []int{1, 1, 0, 1, 1}
	ht[3] = []int{0, 0, 1, 0, 1}
	ht[4] = []int{1, 0, 0, 0, 0}
	dp := make([][]int, n)
	mod := int(1e9) + 7
	for i := range dp {
		dp[i] = make([]int, 5)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				if ht[k][j] == 1 {
					dp[i][j] += dp[i-1][k]
					dp[i][j] %= mod
				}
			}
		}
	}
	res := 0
	for i := 0; i < 5; i++ {
		res += dp[n-1][i]
		res %= mod
	}
	return res
}
