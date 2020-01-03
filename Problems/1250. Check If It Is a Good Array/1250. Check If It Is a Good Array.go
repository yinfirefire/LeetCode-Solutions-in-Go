package main

func isGoodArray(nums []int) bool {
	res := nums[0]
	for i := range nums {
		res = gcd(res, nums[i])
	}
	return res == 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
