package main

func new21Game(N int, K int, W int) float64 {
	//if N is smaller than K, possibility is 0
	//if K-1+W<=N, possibility is 1
	if N < K {
		return float64(0)
	}
	if K+W <= N || K == 0 {
		return float64(1)
	}
	//let dp[i] = possibility of getting i points
	//we can calculate dp[i] from previous step
	// dp[i] = dp[i-1]*P(1) + dp[i-2]*P(2) ... + dp[i-W]*P(W)
	// and P(1) P(2) ... is 1/W
	// so we keep a sliding window which is the sum of last w possibilities
	dp := make([]float64, N+1)
	dp[0] = 1
	wsum := float64(1)
	res := float64(0)
	for i := 1; i <= N; i++ {
		dp[i] = wsum / float64(W)
		if i < K {
			wsum += dp[i]
		} else {
			//once i is greater than k, we cannot draw new numbers
			res += dp[i]
		}
		if i-W >= 0 {
			wsum -= dp[i-W]
		}
	}
	return res
}
