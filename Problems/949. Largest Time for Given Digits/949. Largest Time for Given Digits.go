package main

import "strconv"

//func largestTimeFromDigits(A []int) string {
//	sort.Ints(A)
//	num := 0
//	for i := range A {
//		num = num*10 + A[i]
//	}
//	for i := 23; i >= 0; i-- {
//		for j := 59; j >= 0; j-- {
//			temp := make([]int, 4, 4)
//			temp[0] = i / 10
//			temp[1] = i % 10
//			temp[2] = j / 10
//			temp[3] = j % 10
//			sort.Ints(temp)
//			curr := 0
//			for i := range temp {
//				curr = curr*10 + temp[i]
//			}
//			if curr == num {
//				return fmt.Sprintf("%02d", i) + ":" + fmt.Sprintf("%02d", j)
//			}
//		}
//	}
//	return ""
//}

func largestTimeFromDigits(A []int) string {
	perm := make([]string, 0, 24)
	used := make([]bool, 4, 4)
	dfs(used, "", A, &perm)
	for _, s := range perm {
		h, _ := strconv.Atoi(s[:2])
		m, _ := strconv.Atoi(s[2:])
		if h <= 23 && m <= 59 {
			return s[:2] + ":" + s[2:]
		}
	}
	return ""
}

func dfs(used []bool, curr string, A []int, perm *[]string) {
	if len(curr) == 4 {
		*perm = append(*perm, curr)
		return
	}
	for i := 0; i < 4; i++ {
		if !used[i] {
			used[i] = true
			s := strconv.Itoa(A[i])
			dfs(used, curr+s, A, perm)
			used[i] = false
		}
	}
}
