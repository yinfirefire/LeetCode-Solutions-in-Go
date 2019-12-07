package main

import (
	"math"
)

func sampleStats(count []int) []float64 {
	res := make([]float64, 5, 5)
	var sum = float64(0)
	var cnt = 0
	var min = float64(math.MaxInt32)
	var max = float64(0)
	maxCount := 0
	mode := -1
	for i := 0; i <= 255; i++ {
		if count[i] != 0 {
			min = math.Min(min, float64(i))
			max = math.Max(max, float64(i))
			sum += float64(i * count[i])
			cnt += count[i]
			if count[i] > maxCount {
				mode = i
				maxCount = count[i]
			}
		}
	}
	res[0] = min
	res[1] = max
	res[2] = sum / float64(cnt)
	res[4] = float64(mode)
	half := cnt / 2
	if cnt%2 == 1 {
		halfCnt := 0
		for i := 0; i <= 255; i++ {
			halfCnt += count[i]
			if halfCnt > half {
				res[3] = float64(i)
				break
			}
		}
	} else {
		halfCnt := 0
		l := 0
		r := 0
		for i := 0; i <= 255; i++ {
			if halfCnt < half {
				l = i
			}
			halfCnt += count[i]
			if halfCnt > half {
				r = i
				break
			}
		}
		res[3] = float64(l+r) / float64(2)
	}
	return res
}
