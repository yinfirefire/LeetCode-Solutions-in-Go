package main

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	//actually the maxSize is redundant here,
	//if a string has frequency f, then its substring has frequency at least f
	cnt := make(map[string]int)
	res := 0
	l := 0
	r := 0
	ht := [26]int{}
	unique := 0
	//maintain a fixed size window in minSize
	for ; r < minSize-1; r++ {
		ht[s[r]-'a']++
		if ht[s[r]-'a'] == 1 {
			unique++
		}
	}
	for r < len(s) {
		ht[s[r]-'a']++
		if ht[s[r]-'a'] == 1 {
			unique++
		}
		if unique <= maxLetters {
			str := s[l : r+1]
			cnt[str]++
			if cnt[str] > res {
				res = cnt[str]
			}
		}
		ht[s[l]-'a']--
		if ht[s[l]-'a'] == 0 {
			unique--
		}
		l++
		r++
	}
	return res
}
