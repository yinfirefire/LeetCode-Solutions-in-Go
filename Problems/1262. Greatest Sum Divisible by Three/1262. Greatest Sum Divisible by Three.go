package main

func maxSumDivThree(nums []int) int {
	return maxSumDivK(nums, 3)
}

func maxSumDivK(nums []int, k int) int {
	if k == 0 {
		return -1
	}
	dp := make([]int, k)
	//the max sum with mod 0 to k-1
	for _, v := range nums {
		//to prevent repeat addition of the same number
		temp := make([]int, 0)
		temp = append(temp, dp...)
		for i := range dp {
			n := v + temp[i]
			dp[n%k] = max(n, dp[n%k])
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
