package main

func mirrorReflection(p int, q int) int {
	//To see the whole reflection pathway, extend the room
	//the pathway must satisfy the equation m*p = n*q
	//if m is odd and n is odd, hit receptor 1
	//if m is even and n is odd, hit receptor 0
	//if m is odd and n is even, hit receptor 2
	m := q
	n := p
	for m%2 == 0 && n%2 == 0 {
		m /= 2
		n /= 2
	}
	if m%2 == 0 {
		return 0
	}
	if n%2 == 0 {
		return 2
	}
	return 1
}
