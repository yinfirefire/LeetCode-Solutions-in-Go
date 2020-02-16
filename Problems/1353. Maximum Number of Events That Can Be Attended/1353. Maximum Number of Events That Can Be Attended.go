package main

import (
	"container/heap"
	"sort"
)

type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(a, b int) bool {
	return m[a] < m[b]
}

func (m minHeap) Swap(a, b int) {
	m[a], m[b] = m[b], m[a]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	x := (*m)[len(*m)-1]
	*m = (*m)[0 : len(*m)-1]
	return x
}

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	m := &minHeap{}
	heap.Init(m)
	res := 0
	idx := 0
	l := len(events)
	for i := 1; i <= 100000; i++ {
		for idx < l && events[idx][0] == i {
			heap.Push(m, events[idx][1])
		}
		for m.Len() > 0 && (*m)[0] < i {
			heap.Pop(m)
		}
		if m.Len() > 0 {
			heap.Pop(m)
			res++
		}
	}
	return res
}
