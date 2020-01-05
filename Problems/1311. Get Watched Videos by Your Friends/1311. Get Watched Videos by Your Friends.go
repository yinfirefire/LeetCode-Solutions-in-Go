package main

import "sort"

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	ht := make(map[string]int, 0)
	set := make(map[int]bool, 0)
	queue := make([]int, 0, 0)
	set[id] = true
	queue = append(queue, id)
	for level >= 1 {
		level--
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			for _, f := range friends[cur] {
				if !set[f] {
					set[f] = true
					queue = append(queue, f)
				}
			}
		}
		queue = queue[size:]
	}
	for i := 0; i < len(queue); i++ {
		cur := queue[i]
		for _, v := range watchedVideos[cur] {
			ht[v]++
		}
	}
	res := make([]string, 0)
	for k := range ht {
		res = append(res, k)
	}
	sort.Slice(res, func(a, b int) bool {
		if ht[res[a]] != ht[res[b]] {
			return ht[res[a]] < ht[res[b]]
		} else {
			return res[a] < res[b]
		}
	})
	return res
}
