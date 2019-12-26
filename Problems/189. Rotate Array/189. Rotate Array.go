package main

func rotate(nums []int, k int) {
	reverse(0, len(nums)-1, &nums)
	k %= len(nums)
	reverse(0, k-1, &nums)
	reverse(k, len(nums)-1, &nums)
}

func reverse(s, e int, nums *[]int) {
	for s < e {
		(*nums)[s], (*nums)[e] = (*nums)[e], (*nums)[s]
		s++
		e--
	}
}
