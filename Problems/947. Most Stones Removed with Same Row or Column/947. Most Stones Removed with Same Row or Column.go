package main

func removeStones(stones [][]int) int {
	xset := make(map[int]bool)
	yset := make(map[int]bool)
	islands := 0
	for i := range stones {
		x := stones[i][0]
		y := stones[i][1]
		if !xset[x] && !yset[y] {

			islands += 1
		}
		xset[x] = true
		yset[y] = true
	}
	return len(stones) - islands
}
