package main

import (
	. "LeetCode-GoSol/Utils"
)

func mergeKLists(lists []*ListNode) *ListNode {
	end := len(lists) - 1
	if end == -1 {
		return nil
	} else {
		return multiMerge(lists, 0, end)
	}
}

func multiMerge(lists []*ListNode, start int, end int) *ListNode {
	if end == start {
		return lists[start]
	}
	mid := (end + start) / 2
	left := multiMerge(lists, start, mid)
	right := multiMerge(lists, mid+1, end)
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
}
