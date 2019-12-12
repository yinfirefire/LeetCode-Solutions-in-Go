package main

func isHappy(n int) bool {
	set := make(map[int]bool)
	for _, ok := set[n]; !ok; _, ok = set[n] {
		set[n] = true
		if n == 1 {
			return true
		}
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
	}
	return false
}
