package main

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	swapL := -1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			swapL = i - 1
			break
		}
	}
	if swapL == -1 {
		reverse(nums, 0, len(nums)-1)
	} else {
		swapR := swapL + 1
		for i := swapL + 1; i < len(nums); i++ {
			if nums[i] > nums[swapL] && nums[i] <= nums[swapR] {
				swapR = i
			}
		}
		nums[swapL], nums[swapR] = nums[swapR], nums[swapL]
		reverse(nums, swapL+1, len(nums)-1)
	}
}

func reverse(nums []int, start, end int) {
	i := start
	j := end
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		j--
		i++
	}
}
