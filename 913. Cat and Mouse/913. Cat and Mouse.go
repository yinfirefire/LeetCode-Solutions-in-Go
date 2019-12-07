package main

func catMouseGame(graph [][]int) int {
	//bfs from the final status (e.g. cat/mouse win, draw)
	//for each known status, get all its parent status,
	//for each parent status, check if all the child status is the same
	//if all child status are same, then the parent status is fixed
	n := len(graph)
	queue := make([][]int, 0, 10)
	states := [50][50][3]int{}
	for i := 0; i < n; i++ {
		for j := 1; j <= 2; j++ {
			states[0][i][j] = 1
			queue = append(queue, []int{0, i, j})

			if i != 0 {
				states[i][i][j] = 2
				queue = append(queue, []int{i, i, j})
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		m := cur[0]
		c := cur[1]
		t := cur[2]
		res := states[m][c][t]
		queue = queue[1:]
		for _, p := range getParents(graph, m, c, t) {
			if states[p[0]][p[1]][p[2]] != 0 {
				continue
			}
			if p[2] == res {
				//the parent move is the final winner's move
				//then the final winner can win
				states[p[0]][p[1]][p[2]] = res
				queue = append(queue, p)
			} else if allKidsWin(graph, p[0], p[1], p[2], &states) {
				if p[2] == 1 {
					states[p[0]][p[1]][p[2]] = 2
				} else {
					states[p[0]][p[1]][p[2]] = 1
				}
				queue = append(queue, p)
			}
		}
	}
	return states[1][2][1]
}

func getParents(graph [][]int, m int, c int, turn int) [][]int {
	res := make([][]int, 0, 0)
	if turn == 1 {
		//if current is mouse move
		//then we check cat move for the parents
		for _, next := range graph[c] {
			if next != 0 {
				res = append(res, []int{m, next, 2})
			}
		}
	} else {
		for _, next := range graph[m] {
			res = append(res, []int{next, c, 1})
		}
	}
	return res
}

func allKidsWin(graph [][]int, m int, c int, t int, states *[50][50][3]int) bool {
	//check if all the next step (opponent step leads to win)
	//if all next steps lead to opponent's win, the current state is sure to lose
	if t == 1 {
		//check mouse move
		for _, next := range graph[m] {
			if (*states)[next][c][2] != 2 {
				return false
			}
		}
	} else {
		for _, next := range graph[c] {
			if next == 0 {
				continue
			}
			if (*states)[m][next][1] != 1 {
				return false
			}
		}
	}
	return true
}
