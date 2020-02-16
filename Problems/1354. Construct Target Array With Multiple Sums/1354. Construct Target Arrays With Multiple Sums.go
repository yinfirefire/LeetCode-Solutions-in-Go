package main

import "container/heap"

type maxHeap []int

func (max maxHeap) Len() int {
	return len(max)
}

func (max maxHeap) Less(a, b int) bool {
	return max[a] > max[b]
}

func (max maxHeap) Swap(a, b int) {
	max[a], max[b] = max[b], max[a]
}

func (max *maxHeap) Push(a interface{}) {
	*max = append(*max, a.(int))
}
func (max *maxHeap) Pop() interface{} {
	x := (*max)[len(*max)-1]
	*max = (*max)[0 : len(*max)-1]
	return x
}

func isPossible(target []int) bool {
	m := &maxHeap{}
	heap.Init(m)
	sum := 0
	for i := range target {
		sum += target[i]
		heap.Push(m, target[i])
	}
	for m.Len() > 0 {
		cur := heap.Pop(m).(int)
		if cur == 1 {
			return true
		}
		dif := sum - cur
		times := cur / dif
		if times == 0 {
			return false
		}
		sum -= times * dif
		cur %= dif
		heap.Push(m, cur)
	}
	return true
}
