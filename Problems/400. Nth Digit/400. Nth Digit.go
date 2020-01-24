package main

import (
	"math"
)

func findNthDigit(n int) int {
	//1-9 		-> 9 * length-1 number
	//10-99 	-> 90 * length-2 number
	//100-999 	-> 900 * length-3 number
	//...
	if n <= 9 {
		return n
	}
	nine := 9
	curDigits := 0
	curNum := 0
	i := 1
	for ; i <= 10; i++ {
		curDigits += nine * i
		curNum = curNum*10 + 9
		nine *= 10
		if curDigits+(i+1)*nine > n {
			i++
			break
		}
	}
	dif := (n - curDigits) / i //(difference in digits)/(number of digits) = difference in value
	digit := (n - curDigits) % i
	resNum := curNum + dif
	if digit != 0 {
		//if mod!=0, go to the next number
		resNum += 1
	} else {
		return resNum % 10
	}
	return resNum / int(math.Pow(float64(10), float64(i-digit))) % 10
}
