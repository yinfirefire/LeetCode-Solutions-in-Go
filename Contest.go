package main

import (
	"fmt"
	"sort"
)

func rankTeams(votes []string) string {
	s := votes[0]
	ss := []byte(s)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
	if len(votes) == 1 {
		return votes[0]
	}
	rk := make(map[byte][]int)
	for i := range votes {
		for j := 0; j < len(votes[i]); j++ {
			c := ([]byte(votes[i]))[j]
			if _, ok := rk[c]; !ok {
				rk[c] = make([]int, len(votes[i]))
			}
			rk[c][j]++
		}
	}
	fmt.Println(rk)
	sort.Slice(ss, func(a, b int) bool {
		for i := 0; i < len(votes[0]); i++ {
			if rk[ss[a]][i] != rk[ss[b]][i] {
				return rk[ss[a]][i] > rk[ss[b]][i]
			}
		}
		return ss[a] < ss[b]
	})
	return string(ss)
}

func main() {
	//fmt.Println(rankTeams([]string{"ABC","ACB","ABC","ACB","ACB"}))
	//fmt.Println(rankTeams([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}))
	//fmt.Println(rankTeams([]string{"BCA","CAB","CBA","ABC","ACB","BAC"}))
	fmt.Println(rankTeams([]string{"WXYZ", "XYZW"}))
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
