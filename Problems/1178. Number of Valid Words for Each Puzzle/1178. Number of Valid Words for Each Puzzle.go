package main

func findNumOfValidWords(words []string, puzzles []string) []int {
	mp := make(map[int]int)
	for i := range words {
		c := 0
		for j := range words[i] {
			c |= 1 << uint(words[i][j]-'a')
		}
		mp[c]++
	}
	res := make([]int, len(puzzles))
	for i := range puzzles {
		cp := puzzles[i]
		cnt := 0
		for j := range cp {
			cnt |= 1 << uint(cp[j]-'a')
		}
		sub := cnt
		first := 1 << uint(cp[0]-'a')
		for sub != 0 {
			if sub&first == first {
				// make sure that words contain the first character of current puzzle
				res[i] += mp[sub]
			}
			sub = (sub - 1) & cnt //this will generate every subset of cnt!
		}
	}
	return res
}
