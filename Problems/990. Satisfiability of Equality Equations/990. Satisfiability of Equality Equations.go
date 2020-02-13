package main

func equationsPossible(equations []string) bool {
	root := make([]int, 26)
	for i := range root {
		root[i] = i
	}
	for i := range equations {
		idx1 := int(equations[i][0] - 'a')
		idx2 := int(equations[i][3] - 'a')
		if equations[i][1] == '=' {
			connect(idx1, idx2, root)
		}
	}
	for i := range equations {
		idx1 := int(equations[i][0] - 'a')
		idx2 := int(equations[i][3] - 'a')
		if equations[i][1] != '=' {
			if find(idx1, root) == find(idx2, root) {
				return false
			}
		}
	}
	return true
}

func find(a int, root []int) int {
	if root[a] == a {
		return root[a]
	}
	root[a] = find(root[a], root)
	return root[a]
}

func connect(a, b int, root []int) bool {
	ra := find(a, root)
	rb := find(b, root)
	if ra == rb {
		return false
	} else {
		root[ra] = rb
		return true
	}
}
