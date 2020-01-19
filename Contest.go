package main

import (
	"fmt"
	"sort"
)

func minTaps(n int, ranges []int) int {
	rs := make([][]int, n+1)
	for i := range rs {
		rs[i] = make([]int, 2)
		rs[i][0] = max(0, i-ranges[i])
		rs[i][1] = min(n, i+ranges[i])
	}
	fmt.Println(rs)
	sort.Slice(rs, func(i, j int) bool {
		if rs[i][0] == rs[j][0] {
			return rs[j][1] < rs[i][1]
		} else {
			return rs[i][0] < rs[j][0]
		}
	})
	fmt.Println(rs)
	if rs[0][0] > 0 {
		return -1
	}
	end := rs[0][1]
	cnt := 0
	for i := range rs {
		cnt++
		if end >= n {
			return cnt
		}
		nextEnd := end
		for j := i + 1; j < len(rs); j++ {
			if rs[j][0] <= end && rs[j][1] > nextEnd {
				i = j
				nextEnd = rs[j][1]
			}
		}
		if nextEnd == end {
			return -1
		}
		end = nextEnd
		i--
	}
	return -1
}

func main() {
	fmt.Println(minTaps(9, []int{0, 5, 0, 3, 3, 3, 1, 4, 0, 4}))
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
