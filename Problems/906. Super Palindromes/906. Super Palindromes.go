package main

import (
	"strconv"
)

func superpalindromesInRange(L string, R string) int {
	cands := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 1; i < 10000; i++ {
		cur := strconv.Itoa(i)
		rev := reverse(cur)
		temp, _ := strconv.ParseInt(cur+rev, 10, 64)
		cands = append(cands, temp)
		for j := 0; j < 10; j++ {
			toAdd, _ := strconv.ParseInt(cur+strconv.Itoa(j)+rev, 10, 64)
			cands = append(cands, toAdd)
		}
	}
	res := 0
	l, _ := strconv.ParseInt(L, 10, 64)
	r, _ := strconv.ParseInt(R, 10, 64)
	for _, v := range cands {
		sq := v * v
		if sq >= l && sq <= r && isP(strconv.FormatInt(sq, 10)) {
			res++
		}
	}
	return res
}

func isP(s string) bool {
	chars := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	chars := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
