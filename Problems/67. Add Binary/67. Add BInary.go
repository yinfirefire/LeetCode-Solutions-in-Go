package main

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	res := make([]byte, 0)
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 && a[i] == '1' {
			carry += 1
		}
		if j >= 0 && b[j] == '1' {
			carry += 1
		}
		res = append(res, byte(carry&1)+'0')
		carry >>= 1
		i--
		j--
	}
	reverse(res)
	return string(res)
}

func reverse(arr []byte) []byte {
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}
