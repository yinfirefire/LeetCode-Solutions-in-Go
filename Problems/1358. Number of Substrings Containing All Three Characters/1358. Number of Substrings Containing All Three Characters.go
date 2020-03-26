package main

func numberOfSubstrings(s string) int {
	ca, cb, cc := 0, 0, 0
	i := 0
	j := 0
	res := 0
	for j < len(s) {
		if s[j] == 'a' {
			ca++
		} else if s[j] == 'b' {
			cb++
		} else {
			cc++
		}
		for ca > 0 && cb > 0 && cc > 0 {
			if s[i] == 'a' {
				ca--
			} else if s[i] == 'b' {
				cb--
			} else {
				cc--
			}
			i++
		}
		res += i
		j++
	}
	return res
}
