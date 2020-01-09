package main

import (
	"container/heap"
)

//func findKthLargest(nums []int, k int) int {
//	l := 0
//	r := len(nums) - 1
//	return quickSelect(nums, l, r, k)
//}
//
//func quickSelect(a []int, l, r, k int) int {
//	pivot := l
//	for i := l; i < r; i++ {
//		if a[i] >= a[r] {
//			a[i], a[pivot] = a[pivot], a[i]
//			pivot += 1
//		}
//	}
//	a[pivot], a[r] = a[r], a[pivot]
//	if pivot+1 == k {
//		return a[pivot]
//	} else if pivot+1 > k {
//		return quickSelect(a, l, pivot-1, k)
//	} else {
//		return quickSelect(a, pivot+1, r, k)
//	}
//}

func findKthLargest(nums []int, k int) int {
	pq := &maxHeap{}
	heap.Init(pq)

	for _, v := range nums {
		heap.Push(pq, v)
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	return heap.Pop(pq).(int)
}

type maxHeap []int

func (pq maxHeap) Len() int {
	return len(pq)
}

func (pq maxHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq maxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *maxHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *maxHeap) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}
