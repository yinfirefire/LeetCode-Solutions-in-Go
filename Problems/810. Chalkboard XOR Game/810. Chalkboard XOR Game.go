package main

func xorGame(nums []int) bool {
	//the inverse of xor is still xor
	xor := 0
	for i := range nums {
		xor ^= nums[i]
	}
	//if the xor of all nums is already zero, Alice win
	if xor == 0 {
		return true
	}
	//then check if the number of elements is even, will Alice ever lose?
	//When will Alice lose if the number of elements is even? Alice remove one number and the rest of the xor is 0
	//The situation is s = x1^x2^x3....^xn != 0(n is even)
	//if Alice lose, then remove any x from x1 to xn will result s^xk = 0 (k=1,2,3,4...,n)
	//so (x1^s)^(x2^s)^(x3^s)^(x4^s)....^(xn^s) = 0
	//however if we remove the parentheses, we have
	//x1^x2^x3...^xn^(s^s^s^s...^s) = s^0 = s!=0
	//(there are n s, which is even)
	//which is contradict to the assumption, so Alice will never lose when there is even number
	return len(nums)%2 == 0
}
