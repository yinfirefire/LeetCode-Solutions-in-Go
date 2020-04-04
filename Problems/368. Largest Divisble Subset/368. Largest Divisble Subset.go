package main

import "sort"

type pair struct {
	prev int
	cnt  int
}

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	sort.Ints(nums)
	dp := make([]pair, n)
	dp[0] = pair{-1, 1}
	maxCnt := 1
	maxPos := 0
	for i := 1; i < n; i++ {
		curCnt := 0
		curPrev := -1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if dp[j].cnt > curCnt {
					curCnt = dp[j].cnt
					curPrev = j
				}
			}
		}
		dp[i] = pair{curPrev, curCnt + 1}
		if dp[i].cnt > maxCnt {
			maxCnt = dp[i].cnt
			maxPos = i
		}
	}
	// fmt.Println(dp, maxCnt)
	res := make([]int, maxCnt)
	for maxPos != -1 {
		res[maxCnt-1] = nums[maxPos]
		maxCnt--
		maxPos = dp[maxPos].prev
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
