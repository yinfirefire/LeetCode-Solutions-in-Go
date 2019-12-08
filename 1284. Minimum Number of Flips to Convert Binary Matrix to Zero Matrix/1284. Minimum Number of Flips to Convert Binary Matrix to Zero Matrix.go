package main

import "strconv"

func minFlips(mat [][]int) int {
	//brute force BFS for every status
	//each status has m*n options to flip
	m := len(mat)
	n := len(mat[0])
	queue := make([]string, 0, 0)
	set := make(map[string]bool)
	target := ""
	temp := ""
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp = temp + strconv.Itoa(mat[i][j])
			target = target + "0"
		}
	}
	queue = append(queue, string(temp))
	set[string(temp)] = true
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for t := 0; t < size; t++ {
			cur := queue[t]
			if cur == target {
				return step
			}
			for i := 0; i < len(cur); i++ {
				next := []byte(cur)
				next[i] = flip(next[i])
				x := i / n
				y := i % n
				if x > 0 {
					next[i-n] = flip(next[i-n])
				}
				if x < m-1 {
					next[i+n] = flip(next[i+n])
				}
				if y > 0 {
					next[i-1] = flip(next[i-1])
				}
				if y < n-1 {
					next[i+1] = flip(next[i+1])
				}
				nextstr := string(next)
				if !set[nextstr] {
					set[nextstr] = true
					queue = append(queue, nextstr)
				}
			}
		}
		step++
		queue = queue[size:]
	}
	return -1
}

func flip(a byte) byte {
	if a == '1' {
		return '0'
	} else {
		return '1'
	}
}
