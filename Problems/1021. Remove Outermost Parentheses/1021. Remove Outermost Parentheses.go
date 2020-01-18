package main

func removeOuterParentheses(S string) string {
	res := ""
	cur := 0
	for i := 0; i < len(S); i++ {
		if S[i] == ')' {
			cur--
		}
		if cur != 0 {
			res += S[i : i+1]
		}
		if S[i] == '(' {
			cur++
		}
	}
	return res
}
