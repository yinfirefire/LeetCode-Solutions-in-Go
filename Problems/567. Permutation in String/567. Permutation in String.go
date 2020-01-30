package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	ht := make(map[byte]int)
	cnt := len(s1)
	for i := range s1 {
		ht[s1[i]]++
	}
	l := 0
	r := 0
	for r < len(s1)-1 {
		ht[s2[r]]--
		if ht[s2[r]] >= 0 {
			cnt--
		}
		r++
	}
	for r < len(s2) {
		ht[s2[r]]--
		if ht[s2[r]] >= 0 {
			cnt--
		}
		r++
		if cnt == 0 {
			return true
		}
		ht[s2[l]]++
		if ht[s2[l]] > 0 {
			cnt++
		}
		l++
	}
	return false
}
