package main

func threeEqualParts(A []int) []int {
	//if there are three equal parts, then the zero/one pattern of the number must be the same for the three number
	//first count the total one in the array
	//the number of ones for each number is cnt/=3
	//to get the exact number, traverse from the rightmost part
	//count one to the target number
	if len(A) < 3 {
		return []int{-1, -1}
	}
	numOnes := 0
	for i := range A {
		if A[i] == 1 {
			numOnes++
		}
	}
	if numOnes == 0 {
		return []int{0, 2}
	}
	if numOnes%3 != 0 {
		return []int{-1, -1}
	}
	cnt := numOnes / 3
	r := len(A) - 1
	rCnt := 0
	for rCnt < cnt {
		if A[r] == 1 {
			rCnt++
		}
		r--
	}
	lp := len(A) - r - 1 //pattern length
	//then remove all leading zero for the leftmost part
	l := 0
	for A[l] == 0 {
		l++
	}
	//check if the pattern matches (left and right)
	for i := 0; i < lp; i++ {
		if A[l+i] != A[r+1+i] {
			return []int{-1, -1}
		}
	}
	m := l + lp
	//remove all the leading zero for the middle part
	for A[m] == 0 {
		m++
	}
	//check if the pattern matches (middle and right)
	for i := 0; i < lp; i++ {
		if A[m+i] != A[r+1+i] {
			return []int{-1, -1}
		}
	}
	return []int{l + lp - 1, m + lp}
}
