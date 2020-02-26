package main

func findPaths(m int, n int, N int, i int, j int) int {
	mod := int(1e9 + 7)
	dp := make([][][]int, N+1)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	shift := []int{0, 1, 0, -1, 0}
	dp[0][i][j] = 0
	for i := 1; i <= N; i++ {
		for x := 0; x < m; x++ {
			for y := 0; y < n; y++ {
				for l := 0; l < 4; l++ {
					nx := x + shift[l]
					ny := y + shift[l+1]
					if nx < 0 || ny < 0 || nx >= m || ny >= n {
						dp[i][x][y]++
					} else {
						dp[i][x][y] += dp[i-1][nx][ny]
						dp[i][x][y] %= mod
					}
				}
			}
		}
	}
	return dp[N][i][j] % mod
}
