package main

import (
	"math"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	dividend := float64(numerator)
	divisor := float64(denominator)
	if math.Mod(dividend, divisor) == 0 {
		res := strconv.Itoa(int((dividend / divisor)))
		return res
	}
	sign := 1
	if divisor < 0 {
		divisor = -1 * divisor
		sign *= -1
	}
	if dividend < 0 {
		dividend = -1 * dividend
		sign *= -1
	}
	res := ""
	res = res + strconv.Itoa(int(dividend/divisor))
	fraction := ""
	dividend = math.Mod(dividend, divisor)
	ht := make(map[float64]int)
	pos := 0
	for _, ok := ht[dividend]; !ok && dividend != 0; _, ok = ht[dividend] {
		ht[dividend] = pos
		pos++
		dividend *= 10
		fraction = fraction + strconv.Itoa(int(dividend/divisor))
		dividend = math.Mod(dividend, divisor)
	}
	if dividend == 0 {
		res = res + "." + fraction
	} else {
		pos = ht[dividend]
		res = res + "." + fraction[0:pos] + "(" + fraction[pos:] + ")"
	}
	if sign == -1 {
		res = "-" + res
	}
	return res
}
