package main

func xorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, len(arr)+1)
	prefix[0] = 0
	for i := 1; i <= len(arr); i++ {
		prefix[i] = prefix[i-1] ^ arr[i-1]
	}
	res := make([]int, 0)
	for i := range queries {
		cur := queries[i]
		res = append(res, prefix[cur[1]+1]^prefix[cur[0]])
	}
	return res
}
