package main

func minInsertions(s string) int {
	//let res[i][j] equals to the minimum steps of insertion to make s[i:j] a palindrome

	//for s[i,j]
	//if s[i]==s[j], then res[i][j] = res[i+1][j-1]
	//else res[i][j] = min(res[i+1][j], res[i][j-1])+1
	n := len(s)
	res := twodarray(n, n)
	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			if s[i] == s[j] {
				res[i][j] = res[i+1][j-1]
			} else {
				res[i][j] = min(res[i+1][j], res[i][j-1]) + 1
			}
		}
	}
	return res[0][n-1]
}

func twodarray(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
