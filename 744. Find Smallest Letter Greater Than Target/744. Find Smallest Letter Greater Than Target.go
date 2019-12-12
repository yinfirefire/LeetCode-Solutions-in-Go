package main

func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	r := len(letters) - 1
	for l < r {
		mid := (l + r) / 2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l == len(letters)-1 && letters[l] <= target {
		return letters[0]
	}
	return letters[l]
}
