package main

import (
	"fmt"
	"math"
	"sort"
)

func sumSubseqWidths(A []int) int {
	sort.Ints(A)
	n := len(A)
	fmt.Println(n)
	res := 0
	mod := int(math.Pow(10, 9)) + 7
	for i := 0; i < n; i++ {
		res *= 2
		res += A[n-1-i] - A[i]
		//when A[i] is min, there are 2^(n-i-1) times
		//when A[n-1-i] is max, there are 2^(n-i-1) times
		res %= mod
	}
	return int((res + mod) % mod)
}
