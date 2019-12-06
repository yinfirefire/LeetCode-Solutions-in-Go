package main

import (
	"strconv"
	"unicode"
)

func shortestPathAllKeys(grid []string) int {
	move := [5]int{0, 1, 0, -1, 0}
	sx := -1
	sy := -1
	m := len(grid)
	n := len(grid[0])
	numKeys := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '@' {
				sx = i
				sy = j
			} else if unicode.IsLower(rune(grid[i][j])) {
				numKeys++
			}
		}
	}
	queue := make([][]int, 0, m*n)
	visited := make(map[string]bool)
	cur := strconv.Itoa(0) + " " + strconv.Itoa(sx) + "/" + strconv.Itoa(sy)
	print(cur)
	visited[cur] = true
	queue = append(queue, []int{0, sx, sy})
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			temp := queue[i]
			if temp[0] == (1<<uint(numKeys) - 1) {
				return step
			}
			for j := 0; j < 4; j++ {
				nx := temp[1] + move[j]
				ny := temp[2] + move[j+1]
				keys := temp[0]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] != '#' {
					if unicode.IsLower(rune(grid[nx][ny])) {
						keys |= 1 << uint(grid[nx][ny]-'a')
					}
					if unicode.IsUpper(rune(grid[nx][ny])) {
						if (keys>>uint(grid[nx][ny]-'A'))&1 == 0 {
							continue
						}
					}
					status := strconv.Itoa(keys) + " " + strconv.Itoa(nx) + "/" + strconv.Itoa(ny)
					if !visited[status] {
						visited[status] = true
						queue = append(queue, []int{keys, nx, ny})
					}
				}
			}
		}
		queue = queue[size:]
		step++
	}
	return -1
}
