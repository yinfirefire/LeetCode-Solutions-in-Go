package main

var shift = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func minCost(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}
	queue := make([][]int, 0)
	dfs(grid, res, 0, 0, 0, &queue)
	cost := 0
	for len(queue) > 0 {
		cost += 1
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			x := cur[0]
			y := cur[1]
			for i := 0; i < 4; i++ {
				dfs(grid, res, x+shift[i][0], y+shift[i][1], cost, &queue)
			}
		}
		queue = queue[size:]
	}
	return res[m-1][n-1]
}

func dfs(grid, res [][]int, x, y, cost int, queue *[][]int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || res[x][y] != -1 {
		return
	}
	res[x][y] = cost
	*queue = append(*queue, []int{x, y})
	next := grid[x][y] - 1
	dfs(grid, res, x+shift[next][0], y+shift[next][1], cost, queue)
}
