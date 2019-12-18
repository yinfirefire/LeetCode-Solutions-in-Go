package main

import "strconv"

func maxPoints(points [][]int) int {
	max := 0
	for i, cur := range points {
		curmax := 0
		overlap := 0
		ht := make(map[string]int)

		for j, next := range points {
			if i == j {
				continue
			}
			distX := cur[0] - next[0]
			distY := cur[1] - next[1]
			if distX == 0 && distY == 0 {
				overlap++
				continue
			}
			gcdXY := gcd(distX, distY)
			distX /= gcdXY
			distY /= gcdXY
			cur := strconv.Itoa(distX) + " " + strconv.Itoa(distY)
			ht[cur]++
			if ht[cur] > curmax {
				curmax = ht[cur]
			}
		}
		if max < curmax+overlap+1 {
			max = curmax + overlap + 1
		}
	}
	return max
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
