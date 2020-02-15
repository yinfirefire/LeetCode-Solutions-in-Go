package main

/*
	the key is to find all the cyclic permutations, e.g. 2341 is a cyclic group of 1234, and need 3 swap to restore
	for each 2i, 2i+1, if the person in these two position are not couple, and they come from couple u, v
	so the cyclic group involve couple u and v, which is just like add a connection between couple u and v
*/
var cnt = 0

func minSwapsCouples(row []int) int {
	N := len(row) / 2
	cnt = N
	root := make([]int, cnt)
	for i := range root {
		root[i] = i
	}
	for i := 0; i < N; i++ {
		a := row[i*2]
		b := row[i*2+1]
		connect(root, a/2, b/2)
	}
	return N - cnt
}

func find(root []int, a int) int {
	if root[a] == a {
		return a
	}
	root[a] = find(root, root[a])
	return root[a]
}

func connect(root []int, a, b int) {
	ra := find(root, a)
	rb := find(root, b)
	if ra != rb {
		root[ra] = rb
		cnt--
	}
}
