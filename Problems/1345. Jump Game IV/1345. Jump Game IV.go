package main

func minJumps(arr []int) int {
	ht := make(map[int][]int)
	for i := range arr {
		if ht[arr[i]] == nil {
			ht[arr[i]] = make([]int, 0)
		}
		ht[arr[i]] = append(ht[arr[i]], i)
	}
	numVisited := make(map[int]bool)

	queue := make([]int, 1)
	queue[0] = 0
	step := 0
	visited := make([]bool, len(arr))
	visited[0] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur == len(arr)-1 {
				return step
			}
			/* if the current number has been visited, then all the same number are already in queue */
			if !numVisited[arr[cur]] {
				numVisited[arr[cur]] = true
				for _, j := range ht[arr[cur]] {
					if j == cur {
						continue
					}
					queue = append(queue, j)
					visited[j] = true
				}
			}
			if cur-1 >= 0 && !visited[cur-1] {
				queue = append(queue, cur-1)
				visited[cur-1] = true
			}
			if cur+1 < len(arr) && !visited[cur+1] {
				queue = append(queue, cur+1)
				visited[cur+1] = true
			}
		}
		queue = queue[size:]
		step++
	}
	return -1
}

func main() {
	minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404})
}
