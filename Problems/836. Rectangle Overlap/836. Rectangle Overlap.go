package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	lowx := max(rec1[0], rec2[0])
	lowy := max(rec1[1], rec2[1])
	highx := min(rec1[2], rec2[2])
	highy := min(rec1[3], rec2[3])
	return (lowx < highx) && (lowy < highy)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
