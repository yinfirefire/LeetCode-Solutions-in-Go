package main

func numPrimeArrangements(n int) int {
	mod := int(1e9 + 7)
	cnt := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	np := n - cnt
	res := 1
	for i := 1; i <= np; i++ {
		res *= i
		res %= mod
	}
	for i := 1; i <= cnt; i++ {
		res *= i
		res %= mod
	}
	return res
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
