package main

func grayCode(n int) []int {
	res := make([]int, 0, 0)
	res = append(res, 0)
	for i := 0; i < n; i++ {
		size := len(res)
		for j := size - 1; j >= 0; j-- {
			res = append(res, res[j]|1<<uint(i))
			//bit shift in go must use uint
		}
	}
	return res
}

// n = 1: 0, 1
// n = 2: 00, 01, 11, 10 (or with 1<<1)
// n = 3: 000, 001, 011, 010, 110, 111, 101, 100 (or with 1<<2)
