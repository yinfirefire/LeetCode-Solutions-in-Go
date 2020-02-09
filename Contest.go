package main

func maxStudents(seats [][]byte) int {
	m := len(seats)
	n := len(seats[0])
	res := 0
	helper(0, 0, m, n, 0, seats, &res)
	return res
}

func helper(x, y, m, n, cur int, seats [][]byte, res *int) {
	if x == m-1 && y == n-1 {
		*res = max(*res, cur)
		return
	}
	for i := x; i < m; i++ {
		for j := y; j < n; j++ {
			if seats[i][j] != '#' {
				if j >= 1 && seats[i][j-1] == '1' {
					continue
				}
				if j < n-1 && seats[i][j+1] == '1' {
					continue
				}
				if i >= 1 {
					if j >= 1 && seats[i-1][j-1] == '1' {
						continue
					}
					if j < n-1 && seats[i-1][j+1] == '1' {
						continue
					}
				}
				seats[i][j] = '1'
				if j == n-1 && i != m-1 {
					helper(i+1, 0, m, n, cur+1, seats, res)
				} else {
					helper(i, j+1, m, n, cur+1, seats, res)
				}
				seats[i][j] = '.'
				if j == n-1 && i != m-1 {
					helper(i+1, 0, m, n, cur, seats, res)
				} else {
					helper(i, j+1, m, n, cur, seats, res)
				}
			}
		}
	}
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
