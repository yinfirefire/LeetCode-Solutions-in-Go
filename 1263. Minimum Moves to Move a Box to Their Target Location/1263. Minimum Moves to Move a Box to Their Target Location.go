package main

import (
	"math"
)

func minPushBox(grid [][]byte) int {
	mv := [5]int{0, 1, 0, -1, 0}
	res := math.MaxInt32
	m := len(grid)
	n := len(grid[0])
	queue := make([]int, 0, m*n)
	mem := make(map[int]int)
	target, st, box := [2]int{}, [2]int{}, [2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 'B':
				box = [2]int{i, j}
			case 'T':
				target = [2]int{i, j}
			case 'S':
				st = [2]int{i, j}
			}
		}
	}
	start := encode(box[0], box[1], st[0], st[1])
	queue = append(queue, start)
	mem[start] = 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if mem[cur] > res {
			continue
		}
		pos := decode(cur)
		if pos[0] == target[0] && pos[1] == target[1] {
			if mem[cur] < res {
				res = mem[cur]
			}
			continue
		}
		b := [2]int{pos[0], pos[1]}
		s := [2]int{pos[2], pos[3]}
		for i := 0; i < 4; i++ {
			sx := s[0] + mv[i]
			sy := s[1] + mv[i+1]
			if sx >= 0 && sx < m && sy >= 0 && sy < n && grid[sx][sy] != '#' {
				if sx == b[0] && sy == b[1] {
					nbx := b[0] + mv[i]
					nby := b[1] + mv[i+1]
					if nbx >= 0 && nby >= 0 && nbx < m && nby < n && grid[nbx][nby] != '#' {
						newpos := encode(nbx, nby, sx, sy)
						if val, ok := mem[newpos]; ok && val <= mem[cur]+1 {
							continue
						}
						queue = append(queue, newpos)
						mem[newpos] = mem[cur] + 1
					}
				} else {
					newpos := encode(b[0], b[1], sx, sy)
					if val, ok := mem[newpos]; ok && val <= mem[cur] {
						continue
					}
					mem[newpos] = mem[cur]
					queue = append(queue, newpos)
				}
			}
		}
	}
	if res == math.MaxInt32 {
		return -1
	} else {
		return res
	}
}

func encode(bx int, by int, sx int, sy int) int {
	res := bx<<24 | by<<16 | sx<<8 | sy
	return res
}

func decode(state int) [4]int {
	res := [4]int{}
	res[0] = (state >> 24) & 0xff
	res[1] = (state >> 16) & 0xff
	res[2] = (state >> 8) & 0xff
	res[3] = (state) & 0xff
	return res
}
