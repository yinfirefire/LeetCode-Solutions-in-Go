package main

import "strconv"

func maximumSwap(num int) int {
	str := []byte(strconv.Itoa(num))
	//check from the last digit, update the current max digit
	//if the current digit is smaller than current max, store them as temporary swap position
	curMax := str[len(str)-1]
	maxPos := len(str) - 1
	swapX := -1
	swapY := -1
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] > curMax {
			curMax = str[i]
			maxPos = i
		} else if str[i] < curMax {
			swapX = i
			swapY = maxPos
		}
	}
	if swapX == -1 && swapY == -1 {
		return num
	}
	str[swapX], str[swapY] = str[swapY], str[swapX]
	res, _ := strconv.Atoi(string(str))
	return res
}
