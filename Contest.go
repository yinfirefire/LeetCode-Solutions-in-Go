package main

import (
	"fmt"
	"sort"
	"strings"
)

func largestMultipleOfThree(digits []int) string {
	sum := 0
	sort.Slice(digits, func(a, b int) bool {
		return digits[a] > digits[b]
	})
	if digits[0] == 0 {
		return "0"
	}
	use := make([]int, len(digits))
	for i := range digits {
		sum += digits[i]
		use[i] = 0
	}
	idx := len(digits) - 1
	mod := sum % 3
	cum := 0
	m1 := false
	m2 := false
	if mod == 0 {
		rs := strings.Builder{}
		for i := range digits {
			rs.WriteByte(byte('0' + digits[i]))
		}
		return rs.String()
	}
	for sum%3 != 0 {
		if idx < 0 {
			if !m1 && !m2 {
				return ""
			} else {
				break
			}
		}
		if digits[idx] == 0 {
			idx--
			continue
		}
		if digits[idx]%3 == mod {
			if !m1 {
				use[idx] = 1
				m1 = true
			}
		} else {
			if !m2 {
				if digits[idx]%3 != 0 {
					use[idx] = 2
					cum += digits[idx] % 3
				}
				if cum%3 == mod {
					m2 = true
				}
			}
		}
		idx--
	}
	sb1 := strings.Builder{}
	if m1 {
		c1 := 0
		for i := range digits {
			if use[i] == 0 || use[i] == 2 {
				if use[i] == 2 {
					c1++
				}
				sb1.WriteByte(byte('0' + digits[i]))
			}
		}
		if c1 == len(sb1.String()) {
			sb1 = strings.Builder{}
		}
	}
	sb2 := strings.Builder{}
	if m2 {
		c2 := 0
		for i := range digits {
			if use[i] == 0 || use[i] == 1 {
				if use[i] == 1 {
					c2 += 1
				}
				sb2.WriteByte(byte('0' + digits[i]))
			}
		}
		if c2 == len(sb2.String()) {
			sb2 = strings.Builder{}
		}
	}
	if sb1.Len() > sb2.Len() {
		return sb1.String()
	} else if sb1.Len() < sb2.Len() {
		return sb2.String()
	}
	if strings.Compare(sb1.String(), sb2.String()) > 0 {
		return sb1.String()
	} else {
		return sb2.String()
	}
}

func main() {
	fmt.Print(largestMultipleOfThree([]int{1, 1, 1, 2}))
	fmt.Println(largestMultipleOfThree([]int{5, 8}))
	fmt.Println(largestMultipleOfThree([]int{9, 8, 8, 6, 6}))
	fmt.Println(largestMultipleOfThree([]int{9, 7, 6, 6, 5, 5}))
	fmt.Println(largestMultipleOfThree([]int{8, 6, 7, 1, 0}))

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func twodarray(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
