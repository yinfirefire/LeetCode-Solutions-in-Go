package main

func regionsBySlashes(grid []string) int {
	/*
		 _ _
		|\ /|
		|/ \|
		if the character is blank, all the four parts: top(0), down(3), left(1), right(2) connect each other
		if the character is \, top and right connect, bottom and left connect
		if the character is /, top and left connect, bottom and right connect
	*/
	n := len(grid)
	root := make([]int, n*n*4)
	for i := range root {
		root[i] = i
	}
	cnt := n * n * 4
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				//connect top and previous row down
				connect(4*(i*n+j)+0, ((i-1)*n+j)*n+3, root, &cnt)
			}
			if j > 0 {
				//connect left and previous column right
				connect(4*(i*n+j)+1, (i*n+j-1)*4+2, root, &cnt)
			}
			if grid[i][j] != '/' {
				//connect top right, left bottom
				connect(4*(i*n+j)+0, 4*(i*n+j)+2, root, &cnt)
				connect(4*(i*n+j)+1, 4*(i*n+j)+3, root, &cnt)
			}
			if grid[i][j] != '\\' {
				//connect top left, right bottom
				connect(4*(i*n+j)+0, 4*(i*n+j)+1, root, &cnt)
				connect(4*(i*n+j)+2, 4*(i*n+j)+3, root, &cnt)
			}
		}
	}

	return cnt
}

func find(root []int, a int) int {
	if root[a] == a {
		return a
	}
	root[a] = find(root, root[a])
	return root[a]
}

func connect(a, b int, root []int, cnt *int) bool {
	ra := find(root, a)
	rb := find(root, b)
	if ra != rb {
		root[ra] = rb
		*cnt -= 1
		return true
	}
	return false
}
