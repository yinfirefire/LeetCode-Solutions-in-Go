package main

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 10000000
		}
	}
	for i := range dp {
		dp[i][i] = 0
	}
	for i := range edges {
		dp[edges[i][0]][edges[i][1]] = edges[i][2]
		dp[edges[i][1]][edges[i][0]] = edges[i][2]
	}
	//note! the outer most iteration is the transit point for dynamic programming
	for k := range dp {
		for i := range dp {
			for j := range dp {
				dp[i][j] = min(dp[i][k]+dp[k][j], dp[i][j])
			}
		}
	}
	//now dp is all shortest path for all nodes
	res := make([]int, n)
	for i := range dp {
		for j := range dp[i] {
			if i != j && dp[i][j] <= distanceThreshold {
				res[i]++
			}
		}
	}
	rs := 0
	cur := res[0]
	for i := range res {
		if res[i] <= cur {
			rs = i
			cur = res[i]
		}
	}
	return rs
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
