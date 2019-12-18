package main

func bitwiseComplement(N int) int {
	xor := 1
	for xor < N {
		xor = (xor << 1) + 1
	}
	return N ^ xor
}
