package main

func findDiagonalOrder(nums [][]int) []int {
	m := len(nums)
	total := 0
	for i := range nums {
		total += len(nums[i])
	}
	res := make([]int, total)
	cur := 0
	x := 0
	y := 0
	maxL := len(nums[0])
	maxP := 0
	app := maxL - (m - maxP)

	for i := 0; i < m; i++ {
		x = i
		y = 0
		app = max(app, len(nums[i])-(m-i))
		for j := 0; j < i+1; j++ {
			nx := x - j
			ny := y + j
			if len(nums[nx]) > ny {
				res[cur] = nums[nx][ny]
				cur++
			}
		}
	}
	for i := 0; i < app; i++ {
		x = m - 1
		y = i + 1
		for j := 0; x-j >= 0; j++ {
			nx := x - j
			ny := y + j
			if len(nums[nx]) > ny {
				res[cur] = nums[nx][ny]
				cur++
			}
		}
	}
	return res
}

func main() {
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func twodarray(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
