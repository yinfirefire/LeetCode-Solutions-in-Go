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
	//dp[i][j][k] is the minimum distance for word[i:] with left finger on j, right finger on k
	//dp[i][j][k] = min(dp[i+1][word[i]][k]+distance(word[i], j), dp[i+1][j][word[i]]+distance(word[i],k))
	for i := len(word) - 1; i >= 0; i-- {
		for j := 0; j < 27; j++ {
			for k := 0; k < 27; k++ {
				distL := dist(j, int(word[i]-'A'), tb)
				distR := dist(k, int(word[i]-'A'), tb)
				dp[i][j][k] = min(distL+dp[i+1][int(word[i]-'A')][k], distR+dp[i+1][j][int(word[i]-'A')])
			}
		}
	}
	return dp[0][26][26]

}

func dist(x, y int, tb map[int][]int) int {
	if x == 26 {
		return 0 //the current finger is not on keyboard
	}
	return abs(tb[x][0]-tb[y][0]) + abs(tb[x][1]-tb[y][1])
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
