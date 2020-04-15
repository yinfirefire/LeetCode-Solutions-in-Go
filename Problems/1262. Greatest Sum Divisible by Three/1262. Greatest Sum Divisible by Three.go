package main

func maxSumDivThree(nums []int) int {
	return maxSumDivK(nums, 3)
}

func maxSumDivK(nums []int, k int) int {
	if k == 0 {
		return -1
	}
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	//dp[i][j] is with 0-jth number, the max sum with mod k = i
	m := nums[0] % k
	dp[0][m] = nums[0]
	//dp[0][nums[0]%k] is the base case, in which case there is only one number with mod = nums[0]%k
	for j := 1; j < n; j++ {
		for i := range dp[j-1] {
			n := nums[j] + dp[j-1][i]
			dp[j][n%k] = max(n, dp[j][n%k])      //if we add the current number
			dp[j][i] = max(dp[j][i], dp[j-1][i]) //if we don't add current number
		}
	}
	return dp[len(nums)-1][0]
}

func maxSumDivK2(nums []int, k int) int {
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
