package main

import "fmt"

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	used := make([]bool, n)
	res := 0
	first := -1
	mod := int(1e9 + 7)
	for i := range speed {
		if speed[i]*efficiency[i] > res {
			res = speed[i] * efficiency[i]
			res %= mod
			first = i
		}
	}
	eff := efficiency[first]
	used[first] = true
	temp := res
	cureff := eff
	for i := 1; i < k; i++ {
		add := -1
		for j := range speed {
			if !used[j] {
				cureff = min(eff, efficiency[j])
				newtemp := res/eff*cureff + cureff*speed[j]
				newtemp %= mod
				if newtemp > temp {
					add = j
					temp = newtemp
				}
			}
		}
		if add == -1 {
			return res
		}
		eff = min(eff, efficiency[add])
		res = temp
		used[add] = true
	}
	return temp
}

func main() {
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 3))
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 4))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func twodarray(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
