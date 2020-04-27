package main

import (
	"sort"
)

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	toSort := make([][]int, m*n)
	idx := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			toSort[idx] = []int{i, j}
			idx++
		}
	}
	sort.SliceStable(toSort, func(a, b int) bool {
		if toSort[a][0]+toSort[a][1] == toSort[b][0]+toSort[b][1] {
			if (toSort[a][0]+toSort[a][1])&1 == 1 {
				return a < b
			} else {
				return a > b
			}
		} else {
			return toSort[a][0]+toSort[a][1] < toSort[b][0]+toSort[b][1]
		}
	})
	res := make([]int, n*m)
	for i := range toSort {
		res[i] = matrix[toSort[i][0]][toSort[i][1]]
	}
	return res
}

func findDiagonalOrder2(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	res := make([]int, n*m)
	cnt := 0
	queue := make([][]int, 1)
	queue[0] = []int{0, 0}
	level := 0
	for len(queue) > 0 {
		level++
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			x := cur[0]
			y := cur[1]
			if level%2 == 1 {
				res[cnt+i] = matrix[x][y]
			} else {
				res[cnt+size-1-i] = matrix[x][y]
			}
			if x+1 < m && !visited[x+1][y] {
				visited[x+1][y] = true
				queue = append(queue, []int{x + 1, y})
			}
			if y+1 < n && !visited[x][y+1] {
				visited[x][y+1] = true
				queue = append(queue, []int{x, y + 1})
			}
		}
		queue = queue[size:]
		cnt += size
	}
	return res
}
