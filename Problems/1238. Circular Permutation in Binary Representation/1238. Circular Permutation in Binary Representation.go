package main

func circularPermutation(n int, start int) []int {
	res := grayCode(n)
	i := 0
	for res[i] != start {
		i++
	}
	rotate(i, res)
	return res
}

func grayCode(n int) []int {
	res := make([]int, 0, 0)
	res = append(res, 0)
	for i := 0; i < n; i++ {
		size := len(res)
		for j := size - 1; j >= 0; j-- {
			res = append(res, res[j]|1<<uint(i))
		}
	}
	return res
}

func reverse(s, e int, nums []int) {
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
}

func rotate(k int, nums []int) {
	l := len(nums)
	k %= l
	reverse(0, l-1, nums)
	reverse(0, k-1, nums)
	reverse(k, l-1, nums)
}

func main() {
	circularPermutation(3, 2)
}
