package main

func longestIncreasingPath(matrix [][]int) int {
	shift := []int{0, 1, 0, -1, 0}
	var m = len(matrix)
	if m == 0 {
		return 0
	}
	var n = len(matrix[0])
	if n == 0 {
		return 0
	}
	ind := make([][]int, m)
	for i := range ind {
		ind[i] = make([]int, n)
	}
	queue := make([][]int, 0)
	for i := range ind {
		for j := range ind[i] {
			cur := matrix[i][j]
			for k := 0; k < 4; k++ {
				nx := i + shift[k]
				ny := j + shift[k+1]
				if nx >= 0 && ny >= 0 && nx < m && ny < n && matrix[nx][ny] < cur {
					ind[i][j] += 1
				}
			}
			if ind[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	res := 0
	for len(queue) > 0 {
		res++
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			for j := 0; j < 4; j++ {
				nx := cur[0] + shift[j]
				ny := cur[1] + shift[j+1]
				if nx >= 0 && ny >= 0 && nx < m && ny < n && matrix[cur[0]][cur[1]] < matrix[nx][ny] {
					ind[nx][ny]--
					if ind[nx][ny] == 0 {
						queue = append(queue, []int{nx, ny})
					}
				}

			}
		}
		queue = queue[size:]
	}
	return res
}

var shift = []int{0, 1, 0, -1, 0}

func longestIncreasingPath2(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
	}
	res := 0
	for i := range matrix {
		for j := range matrix[i] {
			res = max(res, helper(i, j, mem, matrix))
		}
	}
	return res
}

func helper(x, y int, mem, matrix [][]int) int {
	if mem[x][y] != 0 {
		return mem[x][y]
	}
	res := 1
	mx := 0
	for i := 0; i < 4; i++ {
		nx := x + shift[i]
		ny := y + shift[i+1]
		if nx >= 0 && ny >= 0 && nx < len(mem) && ny < len(mem[0]) && matrix[x][y] > matrix[nx][ny] {
			mx = max(mx, helper(nx, ny, mem, matrix))
		}
	}
	mem[x][y] = res + mx
	return res + mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
