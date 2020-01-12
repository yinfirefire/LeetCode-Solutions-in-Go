package main

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	res := (C-A)*(D-B) + (G-E)*(H-F)
	dx := min(C, G) - max(A, E)
	dy := min(D, H) - max(B, F)
	if dx > 0 && dy > 0 {
		res -= dx * dy
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
