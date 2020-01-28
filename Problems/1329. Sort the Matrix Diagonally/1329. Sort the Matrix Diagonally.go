package main

import "sort"

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	for i := 0; i < m; i++ {
		temp := make([]int, 0)
		x := i
		y := 0
		for x < m && y < n {
			temp = append(temp, mat[x][y])
			x++
			y++
		}
		sort.Ints(temp)
		idx := 0
		x = i
		y = 0
		for x < m && y < n {
			mat[x][y] = temp[idx]
			idx++
			x++
			y++
		}
	}
	for j := 0; j < n; j++ {
		temp := make([]int, 0)
		x := 0
		y := j
		for x < m && y < n {
			temp = append(temp, mat[x][y])
			x++
			y++
		}
		sort.Ints(temp)
		idx := 0
		x = 0
		y = j
		for x < m && y < n {
			mat[x][y] = temp[idx]
			idx++
			x++
			y++
		}
	}
	return mat
}
