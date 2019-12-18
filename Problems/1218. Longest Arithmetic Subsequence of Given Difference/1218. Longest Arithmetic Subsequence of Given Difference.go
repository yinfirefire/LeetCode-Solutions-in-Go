package main

func longestSubsequence(arr []int, difference int) int {
	ht := make(map[int]int)
	max := 1
	for i := range arr {
		if num, ok := ht[arr[i]-difference]; ok {
			ht[arr[i]] = num + 1
			if ht[arr[i]] > max {
				max = ht[arr[i]]
			}
		} else {
			ht[arr[i]] = 1
		}
	}
	return max
}
