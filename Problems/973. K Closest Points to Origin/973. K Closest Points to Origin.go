package main

import (
	"container/heap"
	"fmt"
)

func kClosest(points [][]int, K int) [][]int {
	pq := &maxHeap{}
	heap.Init(pq)
	for _, v := range points {
		heap.Push(pq, v)
		if pq.Len() > K {
			heap.Pop(pq)
		}
	}
	res := make([][]int, 0)
	for pq.Len() > 0 {
		res = append(res, heap.Pop(pq).([]int))
	}
	return res
}

type maxHeap [][]int

func (pq maxHeap) Len() int {
	return len(pq)
}

func (pq maxHeap) Less(i, j int) bool {
	ii := pq[i][0]*pq[i][0] + pq[i][1]*pq[i][1]
	jj := pq[j][0]*pq[j][0] + pq[j][1]*pq[j][1]
	return ii > jj
}

func (pq maxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *maxHeap) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *maxHeap) Pop() interface{} {
	ptr := *pq
	x := ptr[len(ptr)-1]
	*pq = ptr[0 : len(ptr)-1]
	return x
}

func KClosest2(points [][]int, K int) [][]int {
	quickSelect(points, 0, len(points)-1, K)
	res := make([][]int, 0)
	for i := 0; i < K; i++ {
		res = append(res, points[i])
	}
	return res
}

func quickSelect(a [][]int, l, r, k int) {
	pivot := l
	for i := l; i < r; i++ {
		if (a[i][0]*a[i][0] + a[i][1]*a[i][1]) <= (a[r][0]*a[r][0] + a[r][1]*a[r][1]) {
			a[i], a[pivot] = a[pivot], a[i]
			pivot += 1
		}
	}
	a[pivot], a[r] = a[r], a[pivot]
	if pivot+1 == k {
		return
	} else if pivot+1 > k {
		quickSelect(a, l, pivot-1, k)
	} else {
		quickSelect(a, pivot+1, r, k)
	}
}

func main() {
	fmt.Print(KClosest2([][]int{{6, 10}, {-3, 3}, {-2, 5}, {0, 2}}, 3))
}
