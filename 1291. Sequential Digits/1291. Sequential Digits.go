package main

import "strconv"

func sequentialDigits(low int, high int) []int {
	res := make([]int, 0, 0)
	l := strconv.Itoa(low)
	r := strconv.Itoa(high)
	llen := len(l)
	rlen := len(r)
	pattern := make([][]byte, 0, 0)
	for i := llen; i <= rlen; i++ {
		str := ""
		for j := 0; j < i; j++ {
			str = str + strconv.Itoa(j)
		}
		pattern = append(pattern, []byte(str))
	}
	for _, pat := range pattern {
		for i := 1; i <= 9; i++ {
			var copy []byte
			copy = append(copy, pat...)
			for j := 0; j < len(copy); j++ {
				copy[j] += byte(i)
				if copy[j] == byte(10) {
					break
				}
			}
			tempRes, _ := strconv.Atoi(string(copy))
			if tempRes >= low && tempRes <= high {
				res = append(res, tempRes)
			}
		}
	}
	return res
}
