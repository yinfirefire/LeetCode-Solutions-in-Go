package main

func checkSubarraySum(nums []int, k int) bool {
	mod := make(map[int]int)
	sum := 0
	mod[sum] = -1
	for i := range nums {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}
		if _, ok := mod[sum]; ok {
			if i-mod[sum] > 1 {
				return true
			}
		} else {
			mod[sum] = i
		}
	}
	return false
}
