package main

func minNumberOfFrogs(croakOfFrogs string) int {
	tb := make(map[byte]int, 5)
	prev := map[byte]byte{
		'c': '0',
		'r': 'c',
		'o': 'r',
		'a': 'o',
		'k': 'a',
	}
	tb['0'] = 100000
	tb['c'] = 0
	tb['r'] = 0
	tb['o'] = 0
	tb['a'] = 0
	tb['k'] = 0
	cnt := 0
	res := 0
	for i := range croakOfFrogs {
		cur := croakOfFrogs[i]
		if tb[cur] >= tb[prev[cur]] {
			return -1
		}
		tb[croakOfFrogs[i]]++
		get := true
		if croakOfFrogs[i] == 'c' {
			cnt++
			res = max(res, cnt)
		}
		if cur == 'k' {
			for _, v := range tb {
				if v == 0 {
					get = false
				}
			}
			if get {
				for k := range tb {
					tb[k]--
				}
				cnt--
			}
		}
	}
	for k, v := range tb {
		if v > 0 && k != '0' {
			return -1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
