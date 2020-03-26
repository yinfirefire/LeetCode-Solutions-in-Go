package main

func countOrders(n int) int {
	mod := int(1e9) + 7
	res := 1
	// dp[i] = dp[i-1] * （1+2+3+...+ (i-1)*2+1)
	cur := 1
	for i := 0; i < n; i++ {
		cur = res * (i + 1) * (2*i + 1)
		cur %= mod
		res = cur
	}
	return res
}
