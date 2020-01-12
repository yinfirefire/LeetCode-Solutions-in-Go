package main

func minFlips(a int, b int, c int) int {
	res := 0
	for i := 0; i < 32; i++ {
		b1, b2, b3 := 0, 0, 0
		if (a>>uint(i))&1 == 1 {
			b1 = 1
		}
		if (b>>uint(i))&1 == 1 {
			b2 = 1
		}
		if (c>>uint(i))&1 == 1 {
			b3 = 1
		}
		if b3 == 0 {
			res += b1 + b2
		} else {
			if b1 == 0 && b2 == 0 {
				res += 1
			}
		}
	}
	return res
}
