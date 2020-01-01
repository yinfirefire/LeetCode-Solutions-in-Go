package main

import "math"

func minSwaps(data []int) int {
	cnt := 0
	for _, i := range data {
		if i == 1 {
			cnt++
		}
	}
	l := 0
	r := 0
	curr := 0
	for ; r < cnt-1; r++ {
		if data[r] == 1 {
			curr++
		}
	}
	res := math.MaxInt32
	for r < len(data) {
		if data[r] == 1 {
			curr++
		}
		temp := cnt - curr
		if temp < res {
			res = temp
		}
		if data[l] == 1 {
			curr--
		}
		l++
		r++
	}
	return res
}
