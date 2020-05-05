package main

func numberWays(hats [][]int) int {
	mod := int(1e9) + 7
	m := len(hats)
	htp := make([][]int, 41)
	//htp[i] is the people who like the ith hat
	for i := range hats {
		for j := range hats[i] {
			if htp[hats[i][j]] == nil {
				htp[hats[i][j]] = make([]int, 1)
				htp[hats[i][j]][0] = i
			} else {
				htp[hats[i][j]] = append(htp[hats[i][j]], i)
			}
		}
	}
	dp := make([][]int, 41)
	for i := range dp {
		//use binary to identify the people hats state (e.g. 1010 means the second and fourth people have assigned hats
		dp[i] = make([]int, 1<<uint(m))
	}
	// dp[i][j] means that with already i hats, the number of ways to achieve state j
	dp[0][0] = 1
	for i := 1; i <= 40; i++ {
		for j := 0; j < (1 << uint(m)); j++ {
			dp[i][j] += dp[i-1][j]
			dp[i][j] %= mod //if we don't assign this hat to anyone
			for k := 0; k < len(htp[i-1]); k++ {
				cur := htp[i-1][k]
				//cur is the people we can assign ith hat to
				if (j & (1 << uint(cur))) == 0 {
					//if in jth state, cur people do not have hat assigned
					dp[i][j|(1<<uint(cur))] += dp[i-1][j]
					dp[i][j|(1<<uint(cur))] %= mod
				}
			}
		}
	}
	return dp[40][(1<<uint(m))-1]
}
