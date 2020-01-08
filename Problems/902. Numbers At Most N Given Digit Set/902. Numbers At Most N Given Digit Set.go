package main

import (
	"math"
	"strconv"
)

func atMostNGivenDigitSet(D []string, N int) int {
	s := strconv.Itoa(N)
	n := len(s)
	res := 0
	for i := 1; i < n; i++ {
		res += int(math.Pow(float64(len(D)), float64(i)))
	}
	//first calculate the total number that with fewer digits than N
	//len(D)^(1 ~ len(N)-1)
	ht := make(map[string]bool, 0)
	for _, v := range D {
		ht[v] = true
	}
	i := 0
	for ; i < n; i++ {
		pos := 0
		for j := range D {
			if D[j][0] < s[i] {
				pos++
			}
		}
		res += pos * int(math.Pow(float64(len(D)), float64(n-i-1)))
		if !ht[s[i:i+1]] {
			break
		}
	}
	if i == n {
		return res + 1
	} else {
		return res
	}
}
