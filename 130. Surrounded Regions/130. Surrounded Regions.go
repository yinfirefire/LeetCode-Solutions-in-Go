package main

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			bfs(board, '1', i, 0)
		}
		if board[i][n-1] == 'O' {
			bfs(board, '1', i, n-1)
		}
	}
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			bfs(board, '1', 0, i)
		}
		if board[m-1][i] == 'O' {
			bfs(board, '1', m-1, i)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				bfs(board, 'X', i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}

func bfs(board [][]byte, c byte, x int, y int) {
	m := len(board)
	n := len(board[0])
	move := [5]int{0, 1, 0, -1, 0}
	queue := make([][]int, 0, m*n)
	queue = append(queue, []int{x, y})
	board[x][y] = c
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			nx := cur[0] + move[i]
			ny := cur[1] + move[i+1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && board[nx][ny] == 'O' {
				queue = append(queue, []int{nx, ny})
				board[nx][ny] = c
			}
		}
	}
}
