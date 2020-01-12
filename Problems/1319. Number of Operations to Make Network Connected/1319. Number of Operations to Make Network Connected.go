package main

import "fmt"

func makeConnected(n int, connections [][]int) int {
	root := make([]int, n)
	for i := range root {
		root[i] = i
	}
	dup := 0
	for i := range connections {
		a := connections[i][0]
		b := connections[i][1]
		if !connect(a, b, root) {
			dup++
		}
	}
	set := make(map[int]bool, 0)
	fmt.Println(root)
	for i := range root {
		root[i] = find(i, root)
		set[root[i]] = true
	}
	fmt.Println(root)

	cnt := len(set)
	res := cnt - 1
	if dup < res {
		return -1
	} else {
		return res
	}
}

func find(a int, root []int) int {
	if root[a] == a {
		return a
	}
	root[a] = find(root[a], root)
	return root[a]
}

func connect(a, b int, root []int) bool {
	ra := find(a, root)
	rb := find(b, root)
	if ra == rb {
		return false
	}
	root[rb] = ra
	return true
}
