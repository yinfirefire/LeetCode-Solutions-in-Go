package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backTrack(nums, &res, used, make([]int, 0))
	return res
}

func backTrack(nums []int, res *[][]int, used []bool, temp []int) {
	if len(temp) == len(nums) {
		*res = append(*res, append(make([]int, 0), temp...))
		return
	} else {
		for i := range nums {
			if used[i] {
				continue
			}
			if i > 0 && !used[i-1] && nums[i-1] == nums[i] {
				continue
			}
			used[i] = true
			temp = append(temp, nums[i])
			backTrack(nums, res, used, temp)
			used[i] = false
			temp = temp[0 : len(temp)-1]
		}
	}
}

func main() {
	fmt.Println(permuteUnique([]int{2, 2, 1, 1}))
}
