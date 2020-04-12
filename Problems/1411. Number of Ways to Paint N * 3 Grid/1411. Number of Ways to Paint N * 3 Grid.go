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
