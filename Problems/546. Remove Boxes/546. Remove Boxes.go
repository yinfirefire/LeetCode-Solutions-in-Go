package main

func removeBoxes(boxes []int) int {
	m := len(boxes)
	if m == 0 {
		return 0
	}
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, m)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j <= i; j++ {
			dp[i][i][j] = (j + 1) * (j + 1)
		}
	}
	for j := 1; j < m; j++ {
		for i := j - 1; i >= 0; i-- {
			for k := 0; k <= i; k++ {
				res := (k+1)*(k+1) + dp[i+1][j][0]
				for m := i + 1; m <= j; m++ {
					if boxes[m] == boxes[i] {
						res = max(res, dp[i+1][m-1][0]+dp[m][j][k+1])
					}
				}
				dp[i][j][k] = res
			}
		}
	}
	return dp[0][m-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
