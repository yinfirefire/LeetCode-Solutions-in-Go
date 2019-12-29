package main

func isSolvable(words []string, result string) bool {
	cnt := 0
	ht := make([]int, 26, 26)
	for i := range ht {
		ht[i] = -1
	}
	for _, w := range words {
		chars := []byte(w)
		for i := range chars {
			if ht[chars[i]-'A'] == -1 {
				ht[chars[i]] = cnt
				cnt++
			}
		}
	}
	chars := []byte(result)
	for i := range chars {
		if ht[chars[i]-'A'] == -1 {
			ht[chars[i]] = cnt
			cnt++
		}
	}
	return helper(words, result, make([]int, cnt, cnt), 0, ht, make(map[int]bool))
}

func helper(words []string, result string, value []int, idx int, ht []int, set map[int]bool) bool {
	if idx == len(value) {
		left := 0
		right := 0
		for i := range words {
			chars := []byte(words[i])
			temp := 0
			for j := range chars {
				temp = temp*10 + value[ht[chars[j]-'A']]
			}
			left += temp
		}
		chars := []byte(result)
		for j := range chars {
			right = right*10 + value[ht[chars[j]-'A']]
		}
		return left == right
	} else {
		for i := 0; i <= 9; i++ {
			if !set[i] {
				set[i] = true
				value[idx] = i
				if helper(words, result, value, idx+1, ht, set) {
					return true
				}
			}
			set[i] = false
		}
		return false
	}
}
