package main

import . "LeetCode-GoSol/Utils"

func sumEvenGrandparent(root *TreeNode) int {
	res := 0
	add(false, false, root, &res)
	return res
}

func add(grandpa, pa bool, root *TreeNode, res *int) {
	if root == nil {
		return
	}
	if grandpa {
		*res += root.Val
	}
	if root.Val%2 == 0 {
		add(pa, true, root.Left, res)
		add(pa, true, root.Right, res)
	} else {
		add(pa, false, root.Left, res)
		add(pa, false, root.Right, res)
	}
}
