package main

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	if x+y < z {
		return false
	}
	return z%gcd(x, y) == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
