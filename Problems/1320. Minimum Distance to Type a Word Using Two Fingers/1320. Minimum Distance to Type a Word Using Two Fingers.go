package main

import (
	"math"
	"strconv"
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
	mem := make(map[string]int)
	return helper(word, 0, tb, mem, 'a', 'a')
}

func helper(s string, idx int, tb map[int][]int, mem map[string]int, left, right byte) int {
	if idx == len(s) {
		return 0
	}
	key := strconv.Itoa(idx) + " " + string(left) + " " + string(right)
	if v, ok := mem[key]; ok {
		return v
	}
	res := math.MaxInt32
	c := s[idx]
	x := tb[int(c-'A')][0]
	y := tb[int(c-'A')][1]
	if left == 'a' {
		res = min(res, helper(s, idx+1, tb, mem, c, right))
	} else {
		prevX := tb[int(left-'A')][0]
		prevY := tb[int(left-'A')][1]
		dist := abs(prevX-x) + abs(prevY-y)
		res = min(res, dist+helper(s, idx+1, tb, mem, c, right))
	}
	if right == 'a' {
		res = min(res, helper(s, idx+1, tb, mem, left, c))
	} else {
		prevX := tb[int(right-'A')][0]
		prevY := tb[int(right-'A')][1]
		dist := abs(prevX-x) + abs(prevY-y)
		res = min(res, dist+helper(s, idx+1, tb, mem, left, c))
	}
	mem[key] = res
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

//--------------------------Bottom Up---------------------------------

func minimumDistance2(word string) int {
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
