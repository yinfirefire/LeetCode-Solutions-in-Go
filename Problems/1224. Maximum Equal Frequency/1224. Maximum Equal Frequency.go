package main

//there are only two situations to update the result
//1. [1,1,2,2,3], when i=3, frq(2)=2, 2*2==i+1, then update result to max(result, i+2)
//	means remove 3
//2. [1,1,1,2,2,3], when i=4, frq(2)=2, 2*2==i, then update result to max(result, i+1)
//	means remove 1
func maxEqualFreq(nums []int) int {
	cnt := make(map[int]int) //count of a number
	frq := make(map[int]int) //frequency of the counts
	res := 0
	for i := range nums {
		cnt[nums[i]] = cnt[nums[i]] + 1
		frq[cnt[nums[i]]] = frq[cnt[nums[i]]] + 1
		total := frq[cnt[nums[i]]] * cnt[nums[i]]
		if i == total {
			res = max(res, i+1)
		} else if total == i+1 && i != len(nums)-1 {
			//i cannot be the last index
			res = max(res, i+2)
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
