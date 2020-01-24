package main

func maxRotateFunction(A []int) int {
	if len(A) == 0 {
		return 0
	}
	// F[0] = 0A + 1B + 2C + 3D
	// F[1] = 1A + 2B + 3C + 0D
	// F[2] = 2A + 3B + 0C + 1D
	// F[3] = 3A + 0B + 1C + 2D
	// F[0] -> F[1] = F[0]+(A+B+C+D)-4D
	// F[1] -> F[2] = F[1]+(A+B+C+D)-4C
	sum := 0
	f0 := 0
	for i := range A {
		sum += A[i]
		f0 += i * A[i]
	}
	l := len(A)
	max := f0
	for i := 1; i < l; i++ {
		f0 += sum - l*A[l-i]
		if max < f0 {
			max = f0
		}
	}
	return max
}
