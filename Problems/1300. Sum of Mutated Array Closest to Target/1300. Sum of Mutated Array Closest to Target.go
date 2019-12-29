package main

import "math"

func findBestValue(arr []int, target int) int {
	right := 0
	for i := range arr {
		if arr[i] > right {
			right = arr[i]
		}
	}
	right += 1
	dif := math.MaxUint16
	res := arr[0]
	left := 0
	for left < right {
		mid := (left + right) / 2
		sum := 0
		for _, val := range arr {
			if val > mid {
				sum += mid
			} else {
				sum += val
			}
		}
		temp := int(math.Abs(float64(sum - target)))
		if temp < dif {
			dif = temp
			res = mid
		} else if temp == dif {
			if res > mid {
				res = mid
			}
		}
		if sum < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return res
}
