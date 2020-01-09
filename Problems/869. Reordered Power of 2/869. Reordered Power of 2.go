package main

import "math"

//
//func reorderedPowerOf2(N int) bool {
//	n := strconv.Itoa(N)
//	ct := []byte("0000000000")
//	for i := range n {
//		ct[n[i]-'0']++
//	}
//	set:=make(map[int]bool, 0)
//	l := len(n)
//	upp := 0
//	for i := 0; i < l; i++ {
//		upp *= 10
//		upp += 9
//	}
//	low := int(math.Pow(10, float64(l-1)))
//	pow := 1
//	for pow < upp {
//		if pow >= low {
//			set[pow] = true
//		}
//		pow<<=1
//	}
//	for k:=range set{
//		tempct:=[]byte("0000000000")
//		s:=strconv.Itoa(k)
//		for i:=range s{
//			tempct[s[i]-'0']++
//		}
//		if string(ct)==string(tempct){
//			return true
//		}
//	}
//	return false
//}
//

func reorderedPowerOf2(N int) bool {
	res := count(N)
	cur := 1
	for cur <= math.MaxInt32 {
		if count(cur) == res {
			return true
		} else {
			cur <<= 1
		}
	}
	return false
}

func count(N int) int {
	//count the numbers of 0123456789
	res := 0
	for N > 0 {
		res += int(math.Pow(10, float64(N%10)))
		N /= 10
	}
	return res
}
