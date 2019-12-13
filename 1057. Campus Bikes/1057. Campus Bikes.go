package main

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {

}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[i], pq[j]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return x
}

func assignBikes(workers [][]int, bikes [][]int) []int {

}
