package main

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	d12 := distSq(p1, p2)
	d13 := distSq(p1, p3)
	d14 := distSq(p1, p4)
	if d12 == 0 || d13 == 0 || d14 == 0 {
		return false
	}
	if d12 == d13 && d12*2 == d14 && p4[0] == p3[0]-p1[0]+p2[0] && p4[1] == p3[1]-p1[1]+p2[1] {
		return true
	} else if d12 == d14 && d12*2 == d13 && p3[0] == p2[0]-p1[0]+p4[0] && p3[1] == p2[1]-p1[1]+p4[1] {
		return true
	} else if d13 == d14 && d13*2 == d12 && p2[0] == p3[0]-p1[0]+p4[0] && p2[1] == p3[1]-p1[1]+p4[1] {
		return true
	} else {
		return false
	}
}

func distSq(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}
