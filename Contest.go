package main

import (
	"fmt"
	"sort"
)

func maxJumps(arr []int, d int) int {
	res := 1
	l := len(arr)
	pairs := make([][]int, len(arr))
	for i := range arr {
		pairs[i] = []int{arr[i], i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	dp := make([]int, len(arr))
	for i := range pairs {
		idx := pairs[i][1]
		cur := 1
		for j := 1; j <= d; j++ {
			if idx+j < l {
				if arr[idx+j] > arr[idx] {
					break
				}
				cur = max(cur, dp[idx+j]+1)
			}
		}
		for j := 1; j <= d; j++ {
			if idx-j >= 0 {
				if arr[idx-j] > arr[idx] {
					break
				}
				cur = max(cur, dp[idx-j]+1)
			}
		}
		dp[idx] = cur
		res = max(dp[idx], res)
	}
	return res
}

func main() {
	fmt.Println(maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func twodarray(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
