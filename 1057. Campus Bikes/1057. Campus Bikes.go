package main

import (
	"container/heap"
	"math"
)

type Assignment struct {
	worker   int
	bike     int
	distance float64
}
type MinHeap []Assignment

func (pq MinHeap) Len() int {
	return len(pq)
}

func (pq MinHeap) Less(i, j int) bool {
	if pq[i].distance == pq[j].distance {
		if pq[i].worker == pq[j].worker {
			return pq[i].bike < pq[j].bike
		}
		return pq[i].worker < pq[j].worker
	}
	return pq[i].distance < pq[j].distance
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(Assignment))
}

func (pq *MinHeap) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	mh := MinHeap{}
	l := len(workers)
	b := len(bikes)
	for i := 0; i < l; i++ {
		for j := 0; j < b; j++ {
			dist := math.Abs(float64(workers[i][0]-bikes[j][0])) + math.Abs(float64(workers[i][1]-bikes[j][1]))
			cur := Assignment{worker: i, bike: j, distance: dist}
			heap.Push(&mh, cur)
		}
	}
	usedWorkers := make([]bool, l)
	usedBikes := make([]bool, b)
	res := make([]int, l)
	for mh.Len() > 0 {
		top := heap.Pop(&mh).(Assignment)
		if !usedBikes[top.bike] && !usedWorkers[top.worker] {
			res[top.worker] = top.bike
			usedWorkers[top.worker] = true
			usedBikes[top.bike] = true
		}
	}
	return res
}
