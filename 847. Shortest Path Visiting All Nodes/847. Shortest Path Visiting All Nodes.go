package main

func shortestPathLength(graph [][]int) int {
	//use current node and current visited status as the BFS state
	target := 0
	n := len(graph)
	queue := make([][]int, 0, n*n)
	visited := make(map[string]bool)
	for i := 0; i < n; i++ {
		target |= 1 << uint(i)
		status := 1 << uint(i)
		queue = append(queue, []int{status, i})
		strStatus := string(status) + " " + string(i)
		visited[strStatus] = true
	}
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur[0] == target {
				return step
			}
			for _, next := range graph[cur[1]] {
				nextStatus := string(cur[0]|(1<<uint(next))) + " " + string(next)
				if !visited[nextStatus] {
					visited[nextStatus] = true
					queue = append(queue, []int{cur[0] | (1 << uint(next)), next})
				}
			}
		}
		queue = queue[size:]
		step++
	}
	return step
}

//in Golang bit operation, the number of shifting bits must be unsigned int
