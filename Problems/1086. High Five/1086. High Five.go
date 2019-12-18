package main

import "sort"

func highFive(items [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(items, func(a int, b int) bool {
		if items[a][0] != items[b][0] {
			return items[a][0] < items[b][0]
		} else {
			return items[a][1] >= items[b][1]
		}
	})
	prev := 0
	sum := 0
	cnt := 0
	for _, scores := range items {
		if scores[0] != prev {
			cnt = 0
			prev = scores[0]
			sum = 0
		}
		sum += scores[1]
		cnt++
		if cnt == 5 {
			avg := sum / cnt
			res = append(res, []int{scores[0], avg})
		}
	}
	return res
}
