package main

func strWithout3a3b(A int, B int) string {
	chars := make([]byte, 0, A+B)
	a := byte('a')
	b := byte('b')
	if B > A {
		a, b = b, a
		A, B = B, A
	}
	for A > 0 {
		A--
		chars = append(chars, a)
		if A > B {
			chars = append(chars, a)
			A--
		}
		if B > 0 {
			chars = append(chars, b)
			B--
		}
	}
	return string(chars)
}

/*	In each round, add the char with more number, if its number is still larger than another, append again
	If the smaller number is still larger than zero, append it
*/
