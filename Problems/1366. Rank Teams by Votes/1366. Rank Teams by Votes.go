package main

import (
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
