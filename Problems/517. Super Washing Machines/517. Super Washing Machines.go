package main

func findMinMoves(machines []int) int {
	//since each machine can only transfer one dress at each step
	//the strategy is to find for each machine, the sum of dresses that passes left/right through this machine
	//among all the machine, the highest passes through number is also the steps to take
	sum := 0
	for i := range machines {
		sum += machines[i]
	}
	if sum%len(machines) != 0 {
		return -1
	}
	avg := sum / len(machines)
	leftSum := 0
	leftTarget := 0
	rightTarget := sum + avg
	res := 0
	for i := range machines {
		leftSum += machines[i]                  //prefix sum, i included
		rightSum := sum - leftSum + machines[i] //suffix sum, i included
		leftTarget += avg                       //prefix target, i included
		rightTarget -= avg                      //suffix target, i included
		toRight, toLeft := 0, 0
		if leftSum > leftTarget {
			toRight = leftSum - leftTarget
		}
		if rightSum > rightTarget {
			toLeft = rightSum - rightTarget
		}
		res = max(res, toLeft+toRight)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
