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
