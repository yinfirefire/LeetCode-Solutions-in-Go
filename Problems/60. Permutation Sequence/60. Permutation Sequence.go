package main

import (
	"strconv"
)

func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	product := 1
	number := make([]int, n)
	for i := range number {
		factorial[i] = product
		product *= i + 1
		number[i] = i + 1
	}
	rs := 0
	k--
	for i := n - 1; i >= 0; i-- {
		num := k / factorial[i]
		rs = rs*10 + number[num]
		k %= factorial[i]
		number = append(number[0:num], number[num+1:]...)
	}
	return strconv.Itoa(rs)
}
