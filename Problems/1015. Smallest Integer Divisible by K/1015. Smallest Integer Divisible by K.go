package main

func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	//the last digit is one and the multiple of 2 or 5 will never end with 1
	r := 0
	//the possible length must be within K
	//because there are only K-1 remainder, from length 1 to N, there must be a repeat remainder
	//and the remainder start to appear in circle
	//because currR = (prevR*10+1)%K
	for i := 1; i <= K; i++ {
		r = (r*10 + 1) % K
		if r == 0 {
			return i
		}
	}
	return -1
}
