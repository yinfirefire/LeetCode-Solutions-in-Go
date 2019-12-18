package main

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	shift := []int{0, 1, 0, -1, 0}
	queue := make([][]int, 0, 100)
	cnt := 0
	step := -1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				cnt++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	if cnt == 0 {
		return 0
	}
	for len(queue) > 0 {
		step++
		size := len(queue)
		for i := size - 1; i >= 0; i-- {
			cur := queue[i]
			cx := cur[0]
			cy := cur[1]
			for j := 0; j < 4; j++ {
				nx := cx + shift[j]
				ny := cy + shift[j+1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					cnt--
					grid[nx][ny] = 2
					queue = append(queue, []int{nx, ny})
				}
			}
		}
		queue = queue[size:]
	}
	if cnt == 0 {
		return step
	} else {
		return -1
	}
}
