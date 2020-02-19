package main

import "math"

/*
	Divisible by a or b or c, equals to
	divisible by a + divisible by b + divisible by c - divisible by lcm(a, b) - divisible by lcm(a, c) -divisible by lcm(b, c) + divisible by lcm(a,b,c)
*/

func nthUglyNumber(n int, a int, b int, c int) int {
	l := 1
	r := math.MaxInt32
	for l < r {
		m := l + (r-l)/2
		if count(a, b, c, m) < n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func count(a, b, c, num int) int {
	return num/a + num/b + num/c - num/lcm(a, b) - num/lcm(a, c) - num/lcm(b, c) + num/lcm(c, lcm(a, b))
}
