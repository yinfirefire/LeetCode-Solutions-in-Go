package main

import (
	. "LeetCode-GoSol/Utils"
)

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

/*
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return next
}
*/
