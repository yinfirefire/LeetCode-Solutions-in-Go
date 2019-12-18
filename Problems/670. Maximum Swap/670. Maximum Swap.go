package main

import "strconv"

/*
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
*/

func maximumSwap(num int) int {
	table := make(map[byte]int)
	chars := []byte(strconv.Itoa(num))
	for i, c := range chars {
		table[c] = i
	}
	for i := 0; i < len(chars); i++ {
		for c := byte('9'); c > chars[i]; c-- {
			if table[c] > i {
				chars[table[c]], chars[i] = chars[i], chars[table[c]]
				res, _ := strconv.Atoi(string(chars))
				return res
			}
		}
	}
	return num
}
