package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	length := 0
	for i := range nums {
		idx := binarySearch(dp, nums[i], 0, length)
		if idx == length {
			length++
		}
		dp[idx] = nums[i]
	}
	return length
}

func binarySearch(nums []int, target int, l, r int) int {
	for l < r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
