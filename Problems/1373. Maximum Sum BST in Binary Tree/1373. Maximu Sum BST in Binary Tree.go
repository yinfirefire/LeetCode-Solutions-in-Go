package main

import "math"
import . "LeetCode-GoSol/Utils"

type comp struct {
	isBST    bool
	min, max int
	sum      int
}

func maxSumBST(root *TreeNode) int {
	res := 0
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *int) comp {
	if root == nil {
		return comp{true, math.MaxInt32, math.MinInt32, 0}
	}
	left := helper(root.Left, res)
	right := helper(root.Right, res)
	bst := false
	if left.isBST && right.isBST && root.Val < right.min && root.Val > left.max {
		bst = true
	}
	sum := 0
	if bst {
		sum = root.Val + left.sum + right.sum
		*res = max(*res, sum)
		return comp{true, min(left.min, root.Val), max(right.max, root.Val), sum}
	}
	return comp{false, min(left.min, root.Val), max(right.max, root.Val), sum}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
