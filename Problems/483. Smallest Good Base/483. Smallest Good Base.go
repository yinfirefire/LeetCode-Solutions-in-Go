package main

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	num, _ := strconv.ParseInt(n, 10, 64)
	//the target is to get the longest 1111111....
	//and the longest 1... length is 64, the shortest should be 2, since the range of n is [3,10^18]
	x := int64(1)
	for i := 62; i >= 1; i-- {
		if (x << uint(i)) < num {
			//make sure that the 2 base number is smaller than num
			//now e.g. 10000<num, then we check if there is a number 11111 (base from 2 to num^(1/5)) is equal to num
			if val := binarySearch(num, i); val != -1 {
				return strconv.FormatInt(val, 10)
			}
		}
	}
	return strconv.FormatInt(num-1, 10)
}

func binarySearch(num int64, length int) int64 {
	l := int64(2)
	r := int64(math.Pow(float64(num), 1.0/float64(length))) + 1
	//note, we need to convert length to float64, otherwise, 1.0/int64(n)=0 if n>1
	for l < r {
		mid := l + (r-l)/2
		sum := int64(1)
		cur := int64(1)
		for i := 1; i <= length; i++ {
			cur *= mid
			sum += cur
		}
		if sum == num {
			return mid
		}
		if sum > num {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}
