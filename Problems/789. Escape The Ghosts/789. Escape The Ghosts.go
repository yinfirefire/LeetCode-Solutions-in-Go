package main

func escapeGhosts(ghosts [][]int, target []int) bool {
	//if the ghost want to catch you, the best way to go to the target and wait you
	//so as long as the distance between you and target is smaller than the distance between ghost and target, you are fine
	//is that possible that distance(src, target)<distance(ghost, target) but ghost can catch u on your way?
	//no
	dist := abs(target[0]) + abs(target[1])
	for i := range ghosts {
		temp := 0
		temp += abs(target[0] - ghosts[i][0])
		temp += abs(target[1] - ghosts[i][1])
		if temp < dist {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
