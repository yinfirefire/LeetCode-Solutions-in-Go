package main

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	got := make([]bool, len(status), len(status))
	queue := make([]int, 0, 0)
	used := make([]bool, len(status), len(status))
	for _, i := range initialBoxes {
		got[i] = true
		if status[i] == 1 {
			queue = append(queue, i)
			used[i] = true
		}
	}
	res := 0
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		res += candies[cur]
		for _, i := range keys[cur] {
			status[i] = 1
			if got[i] && !used[i] {
				queue = append(queue, i)
				used[i] = true
			}
		}
		for _, i := range containedBoxes[cur] {
			got[i] = true
			if status[i] == 1 && !used[i] {
				queue = append(queue, i)
				used[i] = true
			}
		}
	}
	return res
}
