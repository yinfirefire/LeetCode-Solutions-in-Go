package main

func matrixBlockSum(mat [][]int, K int) [][]int {
	m := len(mat)
	n := len(mat[0])
	fix := twodarray(m+1, n+1)
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			fix[i][j] = fix[i-1][j] + fix[i][j-1] - fix[i-1][j-1] + mat[i-1][j-1]
		}
	}
	// fmt.Println(fix)
	res := twodarray(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			l := max(0, j-K)
			r := min(n-1, j+K)
			upper := max(0, i-K)
			lower := min(i+K, m-1)
			res[i][j] = fix[lower+1][r+1] + fix[upper][l] - fix[upper][r+1] - fix[lower+1][l]
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func twodarray(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
