package main

func addNegabinary(a []int, b []int) []int {
	i := len(a) - 1
	j := len(b) - 1
	res := make([]int, 0)
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 && a[i] == 1 {
			carry += 1
		}
		if j >= 0 && b[j] == 1 {
			carry += 1
		}
		res = append(res, carry&1)
		//if the current digit is 2 then, it equals to add -1 to the next digit
		carry = -(carry >> 1)
		i--
		j--
	}
	reverse(res)
	for len(res) > 1 && res[0] == 0 {
		//there might be leading zero in this situation
		//e.g. [1,0,1] + [1,1,1,1] = [0,0,0,0]
		res = res[1:]
	}
	return res
}

func reverse(arr []int) []int {
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}
