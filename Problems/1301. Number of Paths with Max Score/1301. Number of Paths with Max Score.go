package main

import "math"

func pathsWithMaxScore(board []string) []int {
	mod := int(math.Pow10(9)) + 7
	m := len(board)
	max := make([][]int, m, m)
	for i := range max {
		max[i] = make([]int, m, m)
	}
	cnt := make([][]int, m, m)
	for i := range cnt {
		cnt[i] = make([]int, m, m)
	}
	max[0][0] = 0
	cnt[0][0] = 1
	for i := 1; i < m; i++ {
		if board[i][0] == 'X' || max[i-1][0] == -1 {
			max[i][0] = -1
			cnt[i][0] = 0
		} else {
			max[i][0] = max[i-1][0] + int(board[i][0]-'0')
			cnt[i][0] = 1
		}
	}
	for i := 1; i < m; i++ {
		if board[0][i] == 'X' || max[0][i-1] == -1 {
			max[0][i] = -1
			cnt[0][i] = 0
		} else {
			max[0][i] = max[0][i-1] + int(board[0][i]-'0')
			cnt[0][i] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < m; j++ {
			if board[i][j] == 'X' {
				max[i][j] = -1
				cnt[i][j] = 0
			} else {
				prevMax := getMax(max[i-1][j-1], getMax(max[i-1][j], max[i][j-1]))
				if prevMax == -1 {
					max[i][j] = -1
					cnt[i][j] = 0
				} else {
					max[i][j] = prevMax + int(board[i][j]-'0')
					if prevMax == max[i-1][j] {
						cnt[i][j] += cnt[i-1][j]
						cnt[i][j] %= mod
					}
					if prevMax == max[i-1][j-1] {
						cnt[i][j] += cnt[i-1][j-1]
						cnt[i][j] %= mod
					}
					if prevMax == max[i][j-1] {
						cnt[i][j] += cnt[i][j-1]
						cnt[i][j] %= mod
					}
				}
			}
		}
	}
	if max[m-1][m-1] == -1 {
		return []int{0, 0}
	} else {
		return []int{max[m-1][m-1] - int('S'-'0'), cnt[m-1][m-1]}
	}
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	pathsWithMaxScore([]string{"E11", "XXX", "11S"})
}
