package main

import "math"

type MinStack struct {
	min int
	arr []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{math.MaxInt32, []int{}}
}

func (this *MinStack) Push(x int) {
	if x <= this.min {
		this.arr = append(this.arr, this.min)
		this.min = x
	}
	this.arr = append(this.arr, x)
}

func (this *MinStack) Pop() {
	n := len(this.arr)
	x := this.arr[n-1]
	if x == this.min {
		this.min = this.arr[n-2]
	}
	this.arr = this.arr[0 : n-2]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
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
