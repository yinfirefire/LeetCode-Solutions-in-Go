package main

func countDigitOne(n int) int {
	q := n
	ans := 0
	x := 1
	for q > 0 {
		digit := q % 10
		q /= 10
		ans += q * x
		if digit == 1 {
			ans += n%x + 1
		} else if digit > 1 {
			ans += x
		}
		x *= 10
	}
	return ans
}

/*
if n = xyzdabc
and we are considering the occurrence of one on thousand, it should be:

(1) xyz * 1000                     if d == 0
(2) xyz * 1000 + abc + 1           if d == 1
(3) xyz * 1000 + 1000              if d > 1
*/
