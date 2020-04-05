package main

func numSteps(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}
	cnt := 0
	ss := []byte(s)
	for i := len(ss) - 1; i > 0; {
		if ss[i] == '0' {
			cnt++
			i--
		} else {
			cnt++
			//change '1' to '0'
			for ss[i] == '1' && i > 0 {
				//count how many zeroes in the tail
				cnt++
				i--
			}
			if i == 0 {
				cnt++
				return cnt
			}
			ss[i] = '1'
		}
	}
	return cnt
}
