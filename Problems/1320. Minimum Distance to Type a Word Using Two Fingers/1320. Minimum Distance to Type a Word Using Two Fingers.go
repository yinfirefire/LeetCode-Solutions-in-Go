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
