package main

import . "LeetCode-GoSol/Utils"

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}
	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
