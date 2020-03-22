package main

func longestPrefix(s string) string {
	pmt := getPMT(s)
	n := len(s)
	if n < 2 {
		return ""
	}
	if pmt[n-1] == 0 {
		return ""
	}
	for i := range pmt {
		if pmt[i] == pmt[n-1] {
			return s[0:pmt[i]]
		}
	}
	return ""
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

func main() {
	longestPrefix("babbb")

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func twodarray(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
