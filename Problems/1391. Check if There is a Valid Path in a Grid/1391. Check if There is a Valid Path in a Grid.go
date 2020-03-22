package main

func hasValidPath(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	root := make([]int, m*n)
	for i := range root {
		root[i] = i
	}
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case 1:
				{
					if j > 0 && (grid[i][j-1] == 1 || grid[i][j-1] == 4 || grid[i][j-1] == 6) {
						connect(i*n+j, i*n+j-1, root)
					}
					if j < n-1 && (grid[i][j+1] == 1 || grid[i][j+1] == 3 || grid[i][j+1] == 5) {
						connect(i*n+j, i*n+j+1, root)
					}
				}
			case 2:
				{
					if i > 0 && (grid[i-1][j] == 2 || grid[i-1][j] == 3 || grid[i-1][j] == 4) {
						connect(i*n+j, (i-1)*n+j, root)
					}
					if i < m-1 && (grid[i+1][j] == 2 || grid[i+1][j] == 5 || grid[i+1][j] == 6) {
						connect(i*n+j, i*n+j+n, root)
					}
				}
			case 3:
				{
					if j > 0 && (grid[i][j-1] == 1 || grid[i][j-1] == 4 || grid[i][j-1] == 6) {
						connect(i*n+j, i*n+j-1, root)
					}
					if i < m-1 && (grid[i+1][j] == 2 || grid[i+1][j] == 5 || grid[i+1][j] == 6) {
						connect(i*n+j, i*n+j+n, root)
					}
				}
			case 4:
				{
					if j < n-1 && (grid[i][j+1] == 1 || grid[i][j+1] == 3 || grid[i][j+1] == 5) {
						connect(i*n+j, i*n+j+1, root)
					}
					if i < m-1 && (grid[i+1][j] == 2 || grid[i+1][j] == 5 || grid[i+1][j] == 6) {
						connect(i*n+j, i*n+j+n, root)
					}
				}
			case 5:
				{
					if i > 0 && (grid[i-1][j] == 2 || grid[i-1][j] == 3 || grid[i-1][j] == 4) {
						connect(i*n+j, (i-1)*n+j, root)
					}
					if j > 0 && (grid[i][j-1] == 1 || grid[i][j-1] == 4 || grid[i][j-1] == 6) {
						connect(i*n+j, i*n+j-1, root)
					}
				}
			case 6:
				{
					if i > 0 && (grid[i-1][j] == 2 || grid[i-1][j] == 3 || grid[i-1][j] == 4) {
						connect(i*n+j, (i-1)*n+j, root)
					}
					if j < n-1 && (grid[i][j+1] == 1 || grid[i][j+1] == 3 || grid[i][j+1] == 5) {
						connect(i*n+j, i*n+j+1, root)
					}
				}
			}
		}
	}
	return connect(0, m*n-1, root)
}

func connect(a, b int, root []int) bool {
	ra := find(a, root)
	rb := find(b, root)
	if ra != rb {
		root[ra] = rb
		return false
	}
	return true
}

func find(a int, root []int) int {
	if root[a] == a {
		return a
	}
	root[a] = find(root[a], root)
	return root[a]
}
