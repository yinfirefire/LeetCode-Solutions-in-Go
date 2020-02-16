package main

import (
	"strconv"
)

/*
	Taking each axis position as a node, and a stone connects x and y axis
*/
var cnt = 0

func removeStones(stones [][]int) int {
	cnt = 0
	root := make(map[string]string)
	for i := range stones {
		x := strconv.Itoa(stones[i][0]) + "X"
		y := strconv.Itoa(stones[i][1]) + "Y"
		connect(root, x, y)
	}

	return len(stones) - cnt
}

func find(root map[string]string, a string) string {
	if _, ok := root[a]; !ok {
		cnt++
		root[a] = a
	}
	if root[a] != a {
		root[a] = find(root, root[a])
	}
	return root[a]
}

func connect(root map[string]string, a, b string) {
	ra := find(root, a)
	rb := find(root, b)
	if ra != rb {
		cnt--
		root[ra] = rb
	}
}
