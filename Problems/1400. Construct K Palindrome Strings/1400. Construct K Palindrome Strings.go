package main

// count the number of odd occurrence
// e.g. if there is one odd occurrence character, then the number of palindrome must be at least 1
// the number of odd occurrence character must be no larger than k
// and the length of string must be no smaller than k
func canConstruct(s string, k int) bool {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}
	odd := 0
	for _, v := range cnt {
		if v%2 != 0 {
			odd++
		}
	}
	return k <= len(s) && odd <= k
}
