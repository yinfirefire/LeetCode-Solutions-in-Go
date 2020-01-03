package main

func hasGroupsSizeX(deck []int) bool {
	//note! must get the two different minimum count!!
	cnt := make(map[int]int, 0)
	for _, v := range deck {
		cnt[v]++
	}
	min1 := len(deck)
	min2 := len(deck)
	for _, v := range cnt {
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v != min1 && v < min2 {
			min2 = v
		}
	}
	if min1 == len(deck) || min2 == len(deck) {
		return len(deck) >= 2
	}
	g := gcd(min1, min2)
	if g < 2 {
		return false
	}
	for _, v := range cnt {
		if v%g != 0 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
