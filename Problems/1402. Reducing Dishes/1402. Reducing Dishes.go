package main

import "sort"

func maxSatisfaction(satisfaction []int) int {
	res := 0
	sort.Ints(satisfaction)
	posSum := 0
	posPos := -1
	for i := range satisfaction {
		if satisfaction[i] >= 0 && posPos == -1 {
			posPos = i
		}
		if satisfaction[i] >= 0 {
			posSum += satisfaction[i]
			res += (i - posPos + 1) * satisfaction[i]
		}
	}
	curRes := res
	for i := 1; i <= posPos; i++ {
		minus := 0
		for j := 1; j <= i; j++ {
			minus += (i - j + 1) * satisfaction[posPos-j]
		}
		res = max(res, curRes+minus+i*posSum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
