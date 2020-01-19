package main

import "sort"

func videoStitching(clips [][]int, T int) int {
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[j][1] < clips[i][1]
		} else {
			return clips[i][0] < clips[j][0]
		}
	})
	if clips[0][0] > 0 {
		return -1
	}
	cur := 0
	end := clips[0][1]
	for i := range clips {
		cur++
		if end >= T {
			return cur
		}
		nextEnd := end
		for j := i + 1; j < len(clips); j++ {
			if clips[j][0] <= end && clips[j][1] > nextEnd {
				nextEnd = clips[j][1]
				i = j
			}
		}
		if nextEnd == end {
			return -1
		}
		i--
		end = nextEnd
	}
	return -1
}
