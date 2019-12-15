package main

import "strconv"

func shortestPath(grid [][]int, k int) int {
	move := []int{0, 1, 0, -1, 0}
	queue := make([][]int, 0, 0)
	queue = append(queue, []int{0, 0, k})
	set := make(map[string]bool)
	set["0"+" "+"0"+" "+strconv.Itoa(k)] = true
	step := 0
	m := len(grid)
	n := len(grid[0])
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur[0] == m-1 && cur[1] == n-1 {
				return step
			}
			for i := 0; i < 4; i++ {
				nx := cur[0] + move[i]
				ny := cur[1] + move[i+1]
				nk := cur[2]
				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					if grid[nx][ny] == 1 {
						nk--
					}
					if nk < 0 {
						continue
					}
					status := strconv.Itoa(nx) + " " + strconv.Itoa(ny) + " " + strconv.Itoa(nk)
					if _, ok := set[status]; !ok {
						set[status] = true
						queue = append(queue, []int{nx, ny, nk})
					}
				}
			}
		}
		queue = queue[size:]
		step++
	}
	return -1
}
