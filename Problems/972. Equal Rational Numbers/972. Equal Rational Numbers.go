package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/strings"
	"strconv"
)

func isRationalEqual(S string, T string) bool {
	return convert(S) == convert(T)
}

func convert(s string) float64 {
	i := strings.Index(s, "(")
	if i > 0 {
		base := s[0:i]
		rep := s[i+1 : len(s)-1]
		for j := 0; j < 20; j++ {
			base += rep
		}
		res, _ := strconv.ParseFloat(base, 64)
		return res
	}
	res, _ := strconv.ParseFloat(s, 64)
	return res
}
