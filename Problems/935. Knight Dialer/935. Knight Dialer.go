package main

func knightDialer(N int) int {
	mod := int(1e9) + 7
	dp := make([]int, 10)
	res := 0
	cnt := map[int][]int{}
	cnt[0] = []int{4, 6}
	cnt[1] = []int{6, 8}
	cnt[2] = []int{7, 9}
	cnt[3] = []int{4, 8}
	cnt[4] = []int{3, 9, 0}
	cnt[5] = []int{}
	cnt[6] = []int{1, 7, 0}
	cnt[7] = []int{2, 6}
	cnt[8] = []int{1, 3}
	cnt[9] = []int{2, 4}
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < N; i++ {
		next := make([]int, 10)
		for j := range next {
			for prev := range cnt[j] {
				next[j] += dp[cnt[j][prev]]
				next[j] %= mod
			}
		}
		dp = next
	}
	for i := range dp {
		res += dp[i]
		res %= mod
	}
	return res
}
