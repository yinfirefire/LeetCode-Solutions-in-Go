package main

import "math"

/*
the question is similar to transform target into a x-base number
since the division operator / only returns  rational number so the only use of / is to get x/x = 1 = x^0
target = a*(x^0)+b*(x^1)+c*(x^2).....
a, b, c can also be negative

To get the least operation, we need to find the largest possible exponent of x
exp = log(target)/log(x)
*/

func leastOpsExpressTarget(x int, target int) int {
	exps := int(math.Log(float64(target))/math.Log(float64(x))) + 1
	//exps is the total bits of base x
	mem := make(map[int]map[int]int, 0)
	return dfs(x, target, exps, mem) - 1 //the first number do not have sign
}

func dfs(x int, target int, k int, mem map[int]map[int]int) int {
	if k == 0 {
		return int(target) * 2
		//if this is the last bit
	}
	if v, ok := mem[int(target)][k]; ok {
		return v
	}
	times := target / int(math.Pow(float64(x), float64(k)))
	//check the number of current bit
	//the final solution of current bit is either times or times+1
	//it cannot be times-1, since then it will never fulfil target
	//if use times+2, the target is over by a larger number, which we need to reduce later, so it won't be a optimized solution
	ans1 := int(times)*k + dfs(x, abs(target-times*int(math.Pow(float64(x), float64(k)))), k-1, mem)
	ans2 := int(times+1)*k + dfs(x, abs(target-(times+1)*int(math.Pow(float64(x), float64(k)))), k-1, mem)
	mem[int(target)] = make(map[int]int, 0)
	mem[int(target)][k] = min(ans1, ans2)
	return min(ans1, ans2)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
