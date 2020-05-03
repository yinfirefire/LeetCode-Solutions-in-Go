package main

import (
	"math"
)

func findRotateSteps(ring string, key string) int {
	mp := map[byte][]int{}
	for i := range ring {
		if _, ok := mp[ring[i]]; !ok {
			mp[ring[i]] = []int{i}
		} else {
			mp[ring[i]] = append(mp[ring[i]], i)
		}
	}
	m := len(ring)
	n := len(key)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32 / 2
		}
	}
	for i := range ring {
		if ring[i] == key[0] {
			dp[0][i] = getDist(0, i, m)
			dp[0][i] += 1
		}
	}
	for i := 1; i < n; i++ {
		cur := key[i]
		for _, curIdx := range mp[cur] {
			for _, prevIdx := range mp[key[i-1]] {
				dp[i][curIdx] = min(dp[i][curIdx], dp[i-1][prevIdx]+getDist(curIdx, prevIdx, m))
			}
			dp[i][curIdx] += 1
		}
	}
	res := math.MaxInt32
	for i := range dp[n-1] {
		res = min(res, dp[n-1][i])
	}
	return res
}

func getDist(a, b, n int) int {
	abs := 0
	rev := 0
	if a < b {
		abs = b - a
		rev = a + n - b
	} else {
		rev = b + n - a
		abs = a - b
	}
	return min(rev, abs)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
