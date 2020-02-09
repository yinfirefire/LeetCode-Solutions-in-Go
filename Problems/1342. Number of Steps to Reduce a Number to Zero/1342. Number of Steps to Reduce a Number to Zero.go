package main

func numberOfSteps(num int) int {
	//bit manipulation
	//if the last bit is 0, divide by two, res+=1
	//if the last bit is 1, minus one and then divide by two, res+=2
	//we will have one extra step included 1-1=0, 0/2=0
	res := 0
	for num > 0 {
		if num&1 == 1 {
			res += 2
		} else {
			res += 1
		}
		num >>= uint(1)
	}
	return res - 1
}
