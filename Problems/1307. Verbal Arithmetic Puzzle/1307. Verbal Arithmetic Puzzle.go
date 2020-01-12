package main

import (
	"fmt"
	"math"
	"sort"
)

type charInt struct {
	x byte
	y int
}

var ws []charInt

func pow10(x int) int {
	res := 1
	for i := 0; i < x; i++ {
		res *= 10
	}
	return res
}

func isSolvable(words []string, result string) bool {
	weights := make(map[byte]int)
	set := make(map[byte]bool) //characters that cannot be zeros (no leading zeros)
	for _, s := range words {
		for i := range s {
			weights[s[i]] += pow10(len(s) - 1 - i)
		}
		if len(s) > 1 {
			set[s[0]] = true //no leading zero
		}
	}
	for i := range result {
		weights[result[i]] -= pow10(len(result) - 1 - i)
	}
	if len(result) > 1 {
		set[result[0]] = true //no leading zero
	}

	for k, v := range weights {
		ws = append(ws, charInt{k, v})
	}

	sort.Slice(ws, func(i, j int) bool {
		return math.Abs(float64(ws[i].y)) > math.Abs(float64(ws[j].y))
	})
	fmt.Println(ws)
	n := len(ws)
	sufSumMax := make([]int, n) //the upper bound of ith weight character value
	sufSumMin := make([]int, n) //the lower bound of ith weight character value
	for i := range ws {
		sufPos := make([]int, 0) //positive suffix weights
		sufNeg := make([]int, 0) //negative suffix weights
		for j := i; j < n; j++ {
			if ws[j].y > 0 {
				sufPos = append(sufPos, ws[j].y)
			} else {
				sufNeg = append(sufNeg, ws[j].y)
			}
		}
		sort.Ints(sufPos)
		sort.Ints(sufNeg)
		for j := range sufPos {
			//note the sufPos is increasing
			sufSumMax[i] += sufPos[j] * (10 - len(sufPos) + j) //add from 10-len(sufPos)+j to ..8,9
			sufSumMin[i] += sufPos[j] * (len(sufPos) - j - 1)  //add from len(sufPos)-j-1 to ..1,0
		}
		for j := range sufNeg {
			//the absolute value is decreasing
			sufSumMin[i] += (9 - j) * sufNeg[j]
			sufSumMax[i] += j * sufNeg[j]
		}
	}
	leadingZero := make([]int, n)
	for i := range ws {
		if set[ws[i].x] {
			leadingZero[i] = 1
		} else {
			leadingZero[i] = 0
		}
	}

	return false
}

func dfs(pos, total int, sufSumMax, sufSumMin []int) bool {
	if pos == len(ws) {
		return total == 0
	}
	return false
}

func main() {
	isSolvable([]string{"SEND", "MORE"}, "MONEY")
}
