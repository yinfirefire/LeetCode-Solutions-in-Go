package main

type Node struct {
	val  int
	Next *Node
	Prev *Node
}

type FirstUnique struct {
	mp   map[int]*Node
	head *Node
	tail *Node
}

func Constructor(nums []int) FirstUnique {
	temp := &Node{-1, nil, nil}
	fq := FirstUnique{map[int](*Node){}, temp, temp}
	for _, v := range nums {
		fq.Add(v)
	}
	return fq
}

func (this *FirstUnique) ShowFirstUnique() int {
	if this.head.Next != nil {
		return this.head.Next.val
	} else {
		return -1
	}
}

func (this *FirstUnique) Add(value int) {
	if _, ok := this.mp[value]; ok {
		td := this.mp[value]
		td.Prev.Next = td.Next
		td.Next.Prev = td.Prev
		delete(this.mp, value)
	} else {
		cur := &Node{value, nil, nil}
		this.tail.Next = cur
		cur.Prev = this.tail
		this.tail = this.tail.Next
	}
}

func main() {
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
