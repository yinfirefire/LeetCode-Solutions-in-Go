package main

import "container/list"

func longestSubarray(nums []int, limit int) int {
	res := 1
	n := len(nums)
	mn := list.New()
	mx := list.New()
	mn.PushBack(0)
	mx.PushBack(0)
	l := 0
	r := 1
	for r < n {
		cur := nums[r]
		for mn.Len() > 0 && nums[mn.Back().Value.(int)] >= cur {
			mn.Remove(mn.Back())
		}
		for mx.Len() > 0 && nums[mx.Back().Value.(int)] <= cur {
			mx.Remove(mx.Back())
		}
		mn.PushBack(r)
		mx.PushBack(r)
		for nums[mx.Front().Value.(int)]-nums[mn.Front().Value.(int)] > limit {
			l++
			for mx.Front().Value.(int) < l {
				mx.Remove(mx.Front())
			}
			for mn.Front().Value.(int) < l {
				mn.Remove(mn.Front())
			}
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
