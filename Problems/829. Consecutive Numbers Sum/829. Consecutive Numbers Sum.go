package main

func consecutiveNumbersSum(N int) int {
	//since its a consective array, starting from x, x+1, x+2 ... x+n-1
	//the sum is x*n+n*(n-1)/2 == N
	//x*n = N-n*(n-1)/2
	//iterate n from 1 to n(n-1)/2>N
	//if (N-n(n-1)/2)%n==0 then res++
	res := 0
	for n := 1; n*(n-1)/2 < N; n++ {
		if (N-n*(n-1)/2)%n == 0 {
			res++
		}
	}
	return res
}
