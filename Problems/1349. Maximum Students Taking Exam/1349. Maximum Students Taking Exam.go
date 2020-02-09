package main

func maxStudents(seats [][]byte) int {
	//bitMask, dp[i][mask] is the max student number of first i rows with the last row distribution status as mask
	m := len(seats)
	n := len(seats[0])
	//the valid position for students for each row
	validPos := make([]int, m)
	for i := range seats {
		cur := 0
		for j := range seats[i] {
			cur = cur << uint(1)
			if seats[i][j] == '.' {
				cur += 1
			}
		}
		validPos[i] = cur
	}
	//the number of students for every mask
	validCnt := make([]int, 1<<uint(n))
	validCnt[0] = 0
	validCnt[1] = 1
	for i := 2; i < (1 << uint(n)); i++ {
		validCnt[i] = validCnt[i/2]
		if i%2 == 1 {
			validCnt[i] += 1
		}
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, 1<<uint(n))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		valid := validPos[i-1] //current slot
		for j := 0; j < 1<<uint(n); j++ {
			if (j&valid == j) && ((j>>1)&j == 0) {
				//first ensure that the current mask is a subset of current row valid position
				//second ensure that no student sitting next to each other
				for k := 0; k < 1<<uint(n); k++ {
					//iterate the last rows status
					if (k>>1)&j == 0 && (j>>1)&k == 0 && dp[i-1][k] != -1 {
						//if the upperLeft and upperRight principle is met
						//if the mask is valid for previous row
						dp[i][j] = max(dp[i][j], dp[i-1][k]+validCnt[j])
					}
				}
			}
		}
	}
	res := 0
	for _, v := range dp[m] {
		res = max(res, v)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
