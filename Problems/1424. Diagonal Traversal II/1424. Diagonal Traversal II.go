package main

import "sort"

func findDiagonalOrder(nums [][]int) []int {
	m := len(nums)
	if m == 0 {
		return []int{}
	}
	total := 0
	pos := make([][]int, 0)
	for i := range nums {
		l := len(nums[i])
		total += l
		temp := make([][]int, l)
		for j := 0; j < l; j++ {
			temp[j] = []int{i, j}
		}
		pos = append(pos, temp...)
	}
	sort.SliceStable(pos, func(a, b int) bool {
		if pos[a][0]+pos[a][1] == pos[b][0]+pos[b][1] {
			return a > b
		} else {
			return pos[a][0]+pos[a][1] < pos[b][0]+pos[b][1]
		}
	})
	res := make([]int, total)
	for i := range pos {
		res[i] = nums[pos[i][0]][pos[i][1]]
	}
	return res
}

func findDiagonalOrder2(nums [][]int) []int {
	m := len(nums)
	if m == 0 {
		return []int{}
	}
	visited := make([][]bool, m)
	total := 0
	for i := range nums {
		l := len(nums[i])
		total += l
		visited[i] = make([]bool, l)
	}
	res := make([]int, total)
	queue := make([][]int, 1)
	queue[0] = []int{0, 0}
	visited[0][0] = true
	cnt := 0
	for len(queue) > 0 {
		cur := queue[0]
		x := cur[0]
		y := cur[1]
		res[cnt] = nums[x][y]
		cnt++
		if x+1 < m && len(nums[x+1]) > y && !visited[x+1][y] {
			queue = append(queue, []int{x + 1, y})
			visited[x+1][y] = true
		}
		if y+1 < len(nums[x]) && !visited[x][y+1] {
			queue = append(queue, []int{x, y + 1})
			visited[x][y+1] = true
		}
		queue = queue[1:]
	}
	return res
}
