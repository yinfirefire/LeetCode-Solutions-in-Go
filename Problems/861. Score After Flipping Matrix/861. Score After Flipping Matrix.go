package main

import "math"

//Greedy: make sure all the number in first column is 1, if not toggle that row
//for the second to last column, if the number of 1 is less than 0, toggle that column
func matrixScore(A [][]int) int {
	for i := 0; i < len(A); i++ {
		if A[i][0] == 0 {
			toggle(A, false, i)
		}
	}
	target := len(A) / 2
	if len(A)%2 != 0 {
		target += 1
	}
	for col := 1; col < len(A[0]); col++ {
		cnt := 0
		for i := 0; i < len(A); i++ {
			if A[i][col] == 1 {
				cnt++
			}
		}
		if cnt < target {
			toggle(A, true, col)
		}
	}
	res := 0
	for i := len(A[0]) - 1; i >= 0; i-- {
		for j := 0; j < len(A); j++ {
			res += int(math.Pow(2, float64(len(A[0])-1-i))) * A[j][i]
		}
	}
	return res
}

func toggle(a [][]int, col bool, pos int) {
	if col {
		for i := 0; i < len(a); i++ {
			a[i][pos] ^= 1
		}
	} else {
		for i := 0; i < len(a[0]); i++ {
			a[pos][i] ^= 1
		}
	}
}
