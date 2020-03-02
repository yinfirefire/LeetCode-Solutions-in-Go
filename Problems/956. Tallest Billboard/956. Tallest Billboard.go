package main

func tallestBillboard(rods []int) int {
	sum := 0
	for i := range rods {
		sum += rods[i]
	}
	dp := make([][]int, len(rods)+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	for i := range rods {
		h := rods[i]
		for j := 0; j <= sum; j++ {
			if dp[i][j] == -1 {
				continue
			}
			dp[i+1][j] = max(dp[i][j], dp[i+1][j])                         //not used
			dp[i+1][j+h] = max(dp[i][j], dp[i+1][j+h])                     //put on higher rod
			dp[i+1][abs(j-h)] = max(dp[i+1][abs(j-h)], dp[i][j]+min(h, j)) //put on shorter rod
		}
	}
	return dp[len(rods)][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
