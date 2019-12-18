package main

func findDuplicate(nums []int) int {
	n := len(nums)
	l := 1
	r := n - 1
	for l < r {
		mid := (l + r) / 2
		cnt := 0
		for _, val := range nums {
			if mid >= val {
				cnt++
			}
		}
		if cnt > mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
