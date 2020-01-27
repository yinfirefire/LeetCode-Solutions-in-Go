package main

func checkValidString(s string) bool {
	buffer := 0
	prev := -1
	left := 0
	used := make([]bool, len(s))
	for i := range s {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			left--
		} else {
			buffer++
			prev = i
		}
		if left < 0 {
			if buffer > 0 {
				buffer--
				used[prev] = true
				left++
			} else {
				return false
			}
		}
	}
	//note, the final result of left can be greater or equal to zero
	right := 0
	buffer = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else if s[i] == '(' {
			right--
		} else {
			if used[i] {
				right--
			} else {
				buffer++
			}
		}
		if right < 0 {
			if buffer > 0 {
				buffer--
				right++
			} else {
				return false
			}
		}
	}
	return right == 0
}
