package main

import (
	"fmt"
)

func minimumDistance(word string) int {
	tb := make(map[int][]int, 26)
	for i := 0; i < 26; i++ {
		x := i / 6
		y := i % 6
		tb[i] = make([]int, 2)
		tb[i][0] = x
		tb[i][1] = y
	}
	dp := make([][][]int, len(word)+1)
	for i := range dp {
		dp[i] = make([][]int, 27)
		for j := range dp[i] {
			dp[i][j] = make([]int, 27)
		}
	}

}

func main() {
	fmt.Println(minimumDistance("CAKE"))
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
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
