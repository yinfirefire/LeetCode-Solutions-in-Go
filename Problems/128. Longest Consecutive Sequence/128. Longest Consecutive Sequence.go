package main

func longestConsecutive(nums []int) int {
	cnt := make(map[int]int)
	res := 0
	for i := range nums {
		if _, ok := cnt[nums[i]]; !ok {
			l := cnt[nums[i]-1]
			r := cnt[nums[i]+1]
			sum := l + r + 1
			res = max(res, sum)
			cnt[nums[i]] = sum
			//only the boundary of a consecutive sequence need to be updated
			cnt[nums[i]-l] = sum
			cnt[nums[i]+r] = sum
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
