package main

import (
	"math"
	"strconv"
)

func crackSafe(n int, k int) string {
	set := make(map[string]bool)
	s := ""
	for i := 0; i < n; i++ {
		s += strconv.Itoa(0)
	}
	set[s] = true
	target := int(math.Pow(float64(k), float64(n)))
	helper(&s, set, target, n, k)
	return s
}

func helper(s *string, set map[string]bool, target, n, k int) bool {
	if len(set) == target {
		return true
	}
	//get last n-1 digits in s
	cur := (*s)[len(*s)-n+1:]
	for i := 0; i < k; i++ {
		temp := cur + strconv.Itoa(i)
		if !set[temp] {
			set[temp] = true
			*s = *s + strconv.Itoa(i)
			if helper(s, set, target, n, k) {
				return true
			}
			delete(set, temp)
			*s = (*s)[0 : len(*s)-1]
		}
	}
	return false
}
