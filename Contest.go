package main

import (
	"fmt"
	"reflect"
)

func numTeams(rating []int) int {
	n := len(rating)
	smaller := make([][]int, 3)
	for i := range smaller {
		smaller[i] = make([]int, n)
	}
	larger := make([][]int, 3)
	for i := range larger {
		larger[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		smaller[0][i] = 1
	}
	for i := 0; i < n; i++ {
		larger[0][i] = 1
	}
	for i := 1; i < 3; i++ {
		for j := 1; j < n; j++ {
			for k := i - 1; k < j; k++ {
				if rating[k] < rating[j] {
					smaller[i][j] += smaller[i-1][k]
				}
			}
		}
	}
	for i := 1; i < 3; i++ {
		for j := 1; j < n; j++ {
			for k := i - 1; k < j; k++ {
				if rating[k] > rating[j] {
					larger[i][j] += larger[i-1][k]
				}
			}
		}
	}
	res := 0
	for i := 2; i < n; i++ {
		res += smaller[2][i]
		res += larger[2][i]
	}
	return res
}

func change(a [2]int) {
	a[0] = 1
	a[1] = 1
}
func change2(b []int) {
	b[0] = 1
	b[1] = 1
}

func main() {
	fmt.Println(numTeams([]int{2, 5, 3, 4, 1}))
	fmt.Println(numTeams([]int{1, 2, 3, 4}))
	fmt.Println(reflect.TypeOf([2]int{}))
	cc := [2]int{2, 2}
	dd := []int{2, 2}
	change(cc)
	change2(dd)
	fmt.Println(cc, dd)
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
