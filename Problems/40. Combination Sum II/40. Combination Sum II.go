package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	helper([]int{}, candidates, 0, 0, target, &res)
	return res
}

func helper(temp, cands []int, idx, sum, target int, res *[][]int) {
	if sum == target {
		cp := make([]int, len(temp))
		copy(cp, temp)
		*res = append(*res, cp)
		return
	} else if sum > target {
		return
	} else {
		for i := idx; i < len(cands); i++ {
			if i > idx && cands[i-1] == cands[i] {
				continue
			}
			sum += cands[i]
			temp = append(temp, cands[i])
			helper(temp, cands, i+1, sum, target, res)
			temp = temp[0 : len(temp)-1]
			sum -= cands[i]
		}
	}
}
