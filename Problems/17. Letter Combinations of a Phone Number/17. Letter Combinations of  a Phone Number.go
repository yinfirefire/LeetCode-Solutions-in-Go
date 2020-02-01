package main

import "strconv"

func letterCombinations(digits string) []string {
	ht := make(map[int]string)
	ht[2] = "abc"
	ht[3] = "def"
	ht[4] = "ghi"
	ht[5] = "jkl"
	ht[6] = "mno"
	ht[7] = "pqrs"
	ht[8] = "tuv"
	ht[9] = "wxyz"
	queue := []string{""}
	for len(digits) > 0 {
		cur, _ := strconv.Atoi(digits[0:1])
		digits = digits[1:]
		size := len(queue)
		for i := 0; i < size; i++ {
			s := queue[i]
			for j := range ht[cur] {
				queue = append(queue, s+string(ht[cur][j]))
			}
		}
		queue = queue[size:]
	}
	return queue
}
