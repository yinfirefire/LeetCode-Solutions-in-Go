package main

import (
	"container/heap"
	"sort"
)

type minHeap []int

func (min minHeap) Len() int {
	return len(min)
}

func (min minHeap) Swap(i, j int) {
	min[i], min[j] = min[j], min[i]
}

func (min minHeap) Less(i, j int) bool {
	return min[i] < min[j]
}

func (min *minHeap) Push(x interface{}) {
	*min = append(*min, x.(int))
}

func (min *minHeap) Pop() interface{} {
	x := (*min)[len(*min)-1]
	*min = (*min)[0 : len(*min)-1]
	return x
}

type pair struct {
	s int
	e int
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	mod := int(1e9) + 7
	m := minHeap{}
	heap.Init(&m)
	arr := make([]pair, n)
	for i := range speed {
		arr[i] = pair{speed[i], efficiency[i]}
	}
	sort.Slice(arr, func(a, b int) bool {
		if arr[a].e == arr[b].e {
			return arr[a].s > arr[b].s
		} else {
			return arr[a].e > arr[b].e
		}
	})
	sum := 0
	res := 0
	for i := range arr {
		if m.Len() == k {
			sum -= heap.Pop(&m).(int)
		}
		heap.Push(&m, arr[i].s)
		sum += arr[i].s
		res = max(res, sum*arr[i].e)
	}
	return res % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
