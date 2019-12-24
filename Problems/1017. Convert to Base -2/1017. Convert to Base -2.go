package main

import "strconv"

func baseNeg2(N int) string {
	res := ""
	for N > 0 {
		c := strconv.Itoa(N & 1)
		res = res + string(c)
		N = -(N >> 1)
	}
	if len(res) == 0 {
		return "0"
	} else {
		return reverse(res)
	}
}

func base2(N int) string {
	res := ""
	for N > 0 {
		c := strconv.Itoa(N & 1)
		res = res + string(c)
		N >>= 1
	}
	if len(res) == 0 {
		return "0"
	} else {
		return reverse(res)
	}
}

func reverse(s string) string {
	chars := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
