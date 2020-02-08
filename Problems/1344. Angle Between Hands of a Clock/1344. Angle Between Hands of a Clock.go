package main

import (
	"math"
)

func angleClock(hour int, minutes int) float64 {
	full := float64(360)
	min := full * float64(minutes) / float64(60)
	h := full*float64(hour%12)/float64(12) + full/float64(12)*float64(minutes)/float64(60)
	return math.Min(math.Abs(h-min), 360-math.Abs(h-min))
}
