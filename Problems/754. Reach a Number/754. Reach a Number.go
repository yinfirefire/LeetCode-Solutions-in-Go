package main

import "math"

func reachNumber(target int) int {
	//first we need to find the steps of total addition 1+2+3+...+n
	target = int(math.Abs(float64(target)))
	step := 0
	total := 0
	for total < target {
		step++
		total += step
	}
	if total == target {
		return step
	}
	//if they are not equal, we need to reduce a number i and the total-=2*i
	//so the difference between target and total should be even
	//if they are not even, we just add steps to make the difference even
	dif := total - target
	for dif%2 != 0 {
		step++
		total += step
		dif = total - target
	}
	return step
}
