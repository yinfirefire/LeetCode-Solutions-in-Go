package main

/*
import "container/heap"

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
func nthUglyNumber(n int) int {
	min := &minHeap{}
	heap.Init(min)
	heap.Push(min, 1)
	for i := 1; i < n; i++ {
		cur := heap.Pop(min).(int)
		for min.Len() > 0 && (*min)[0] == cur {
			heap.Pop(min)
		}
		heap.Push(min, cur*2)
		heap.Push(min, cur*3)
		heap.Push(min, cur*5)
	}
	return heap.Pop(min).(int)
}
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
		if dp[i] == dp[p2]*2 {
			p2++
		}
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
