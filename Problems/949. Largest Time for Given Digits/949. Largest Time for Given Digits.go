package main

import (
	"fmt"
	"sort"
)

func largestTimeFromDigits(A []int) string {
	sort.Ints(A)
	num := 0
	for i := range A {
		num = num*10 + A[i]
	}
	for i := 23; i >= 0; i-- {
		for j := 59; j >= 0; j-- {
			temp := make([]int, 4, 4)
			temp[0] = i / 10
			temp[1] = i % 10
			temp[2] = j / 10
			temp[3] = j % 10
			sort.Ints(temp)
			curr := 0
			for i := range temp {
				curr = curr*10 + temp[i]
			}
			if curr == num {
				return fmt.Sprintf("%02d", i) + ":" + fmt.Sprintf("%02d", j)
			}
		}
	}
	return ""
}
