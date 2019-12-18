package main

var res [][]string

func partition(s string) [][]string {
	res = [][]string{}
	helper(0, s, []string{})
	return res
}

func helper(idx int, s string, cur []string) {
	if idx == len(s) {
		res = append(res, append([]string{}, cur...))
		return
	}
	for i := idx; i < len(s); i++ {
		if isPalindrome(s, idx, i) {
			cur = append(cur, s[idx:i+1])
			helper(i+1, s, cur)
			cur = cur[:len(cur)-1]
		}
	}
}

func isPalindrome(s string, i int, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}
