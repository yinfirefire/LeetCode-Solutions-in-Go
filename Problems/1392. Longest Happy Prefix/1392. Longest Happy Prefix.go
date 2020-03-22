package main

func longestPrefix(s string) string {
	pmt := getPMT(s)
	n := len(s)
	if n < 2 {
		return ""
	}
	return s[0:pmt[len(pmt)-1]]
}

func getPMT(s string) []int {
	n := len(s)
	pmt := make([]int, n)
	pmt[0] = 0
	i := 1
	j := 0
	for i < n {
		if s[i] == s[j] {
			pmt[i] = j + 1
			j++
			i++
		} else if j == 0 {
			pmt[i] = 0
			i++
		} else {
			j = pmt[j-1]
		}
	}
	return pmt
}
