package main

import "math"

func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])
	sum := make([][]int, m+1, m+1)
	for i := range sum {
		sum[i] = make([]int, n+1, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i+1][j+1] = sum[i][j+1] + sum[i+1][j] - sum[i][j] + mat[i][j]
		}
	}
	l := 0
	r := int(math.Min(float64(m), float64(n)))
	for l+1 < r {
		mid := (l + r) / 2
		if !inThreshold(&sum, threshold, mid) {
			r = mid
		} else {
			l = mid
		}
	}
	if inThreshold(&sum, threshold, r) {
		return r
	} else {
		return l
	}
}

func inThreshold(sum *[][]int, threshold int, length int) bool {
	for i := 1; i < len(*sum); i++ {
		for j := 1; j < len((*sum)[0]); j++ {
			if i < length || j < length {
				continue
			}
			if (*sum)[i][j]+(*sum)[i-length][j-length]-(*sum)[i-length][j]-(*sum)[i][j-length] <= threshold {
				return true
			}
		}
	}
	return false
}
