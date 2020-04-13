package main

func numOfWays(n int) int {
	// if the previous row has three color, e.g. ABC
	// the next row can be in two or three colors
	// 3color: CAB BCA
	// 2color: BAB BCB

	// if the previous row has two color, e.g. ABA
	// the next row can also in two/three colors
	// 2color: CAC BAB BCB
	// 3color: CAB, BAC

	// so we can get this conclusion
	// if the current row has three color, then it comes from 2*prevTwoColor + 2*prevThreeColor
	// if the current row has two color, then it comes from 2*prevThreeColor + 3*prevTwoColor
	two := make([]int, n)
	three := make([]int, n)
	two[0] = 6
	three[0] = 6
	mod := int(1e9) + 7
	for i := 1; i < n; i++ {
		two[i] = 3*two[i-1] + 2*three[i-1]
		two[i] %= mod
		three[i] = 2*two[i-1] + 2*three[i-1]
		three[i] %= mod
	}
	return (two[n-1] + three[n-1]) % (int(1e9) + 7)
}

func numOfWays2(n int) int {
	mod := int(1e9) + 7
	//for one row there are 3^3 combinations
	//get the validity of these 27 pattern
	valid := make([]bool, 27)
	for i := range valid {
		a := i / 9       //change every 9 pattern
		b := (i / 3) % 3 //change every 3 pattern
		c := i % 3       //change every 1 pattern
		if a != b && b != c {
			valid[i] = true
		}
	}
	//get the valid table for next row
	vv := make([][]bool, 27)
	for i := range vv {
		vv[i] = make([]bool, 27)
		if valid[i] {
			for j := range vv[i] {
				if valid[j] {
					a := i / 9
					b := j / 9
					c := (i / 3) % 3
					d := (j / 3) % 3
					e := i % 3
					f := j % 3
					if a != b && c != d && e != f {
						vv[i][j] = true
					}
				}
			}
		}
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 27)
	}
	for j := range dp[0] {
		if valid[j] {
			dp[0][j] = 1
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < 27; j++ {
			for k := 0; k < 27; k++ {
				if vv[k][j] {
					dp[i][j] += dp[i-1][k]
					dp[i][j] %= mod
				}
			}
		}
	}
	res := 0
	for i := range dp[n-1] {
		res += dp[n-1][i]
		res %= mod
	}
	return res
}
