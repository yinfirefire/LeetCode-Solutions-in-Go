package main

import . "LeetCode-GoSol/Utils"

func plusOne(head *ListNode) *ListNode {
	next := carry(head)
	if next {
		newHead := ListNode{1, head}
		return &newHead
	} else {
		return head
	}
}

func carry(head *ListNode) bool {
	if head == nil {
		return true
	}
	next := carry(head.Next)
	curVal := head.Val
	if next {
		curVal += 1
	}
	if curVal >= 10 {
		head.Val = curVal % 10
		return true
	} else {
		head.Val = curVal
		return false
	}
}
