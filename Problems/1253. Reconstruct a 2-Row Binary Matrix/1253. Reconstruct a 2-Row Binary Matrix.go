package main

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	res := [][]int{make([]int, n), make([]int, n)}
	for i := range colsum {
		if colsum[i] == 2 {
			upper--
			lower--
			res[0][i] = 1
			res[1][i] = 1
		}
	}
	if upper < 0 || lower < 0 {
		return make([][]int, 0)
	}
	for i := range colsum {
		if colsum[i] == 1 {
			if upper > 0 {
				upper--
				res[0][i] = 1
			} else if lower > 0 {
				lower--
				res[1][i] = 1
			} else {
				return make([][]int, 0)
			}
		}
	}
	if upper != 0 || lower != 0 {
		return make([][]int, 0)
	}
	return res
}
