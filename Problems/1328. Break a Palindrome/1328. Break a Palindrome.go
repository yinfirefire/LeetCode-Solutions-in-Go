package main

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}
	pb := []byte(palindrome)
	mid := len(palindrome) / 2
	for i := 0; i < mid; i++ {
		if palindrome[i] > 'a' {
			pb[i] = 'a'
			return string(pb)
		}
	}
	pb[len(pb)-1] = 'b'
	return string(pb)
}
