package main

import "math"

var mem map[string]int

func palindromePartition(s string, k int) int {
	mem = make(map[string]int)
	dp := [100][101]int{}
	for idx := range s {
		dp[0][idx+1] = helper(s[0 : idx+1])
	}

	for i := 1; i < k; i++ {
		for j := i; j <= len(s); j++ {
			cur := math.MaxInt32
			for p := j; p >= i; p-- {
				cur = min(cur, dp[i-1][p-1]+helper(s[p-1:j]))
			}
			dp[i][j] = cur
		}
	}
	return dp[k-1][len(s)]
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func helper(str string) int {
	if len(str) == 0 {
		return 0
	}
	if val, ok := mem[str]; ok {
		return val
	}
	l := 0
	r := len(str)
	res := 0
	for l < r {
		if str[l] != str[r] {
			res++
		}
		l++
		r--
	}
	mem[str] = res
	return res
}
