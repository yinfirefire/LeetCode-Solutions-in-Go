package main

func uniqueLetterString(s string) int {
	//what we need to do is to find the previous and next appearance of the character
	//then we know the longest substring it can be in, so that this character is unique
	//e.g. XAXXAXA the second A is unique in XA(XXAX)A
	n := len(s)
	occur := [26][2]int{}
	for i := range occur {
		occur[i] = [2]int{-1, -1}
	}
	sum := 0
	mod := int(1e9) + 7
	for i := range s {
		cur := s[i] - 'A'
		sum += (i - occur[cur][1]) * (occur[cur][1] - occur[cur][0])
		sum %= mod
		occur[cur] = [2]int{occur[cur][1], i}
	}
	for i := range occur {
		sum += (n - occur[i][1]) * (occur[i][1] - occur[i][0])
		sum %= mod
	}
	return sum
}
