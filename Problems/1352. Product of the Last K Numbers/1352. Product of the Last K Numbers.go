package main

type ProductOfNumbers struct {
	prefixProd []int
	lastZero   int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{[]int{1}, 0}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prefixProd = append(this.prefixProd, 1)
		this.lastZero = len(this.prefixProd) - 1
	} else {
		this.prefixProd = append(this.prefixProd, this.prefixProd[len(this.prefixProd)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	start := len(this.prefixProd) - 1 - k
	if this.lastZero > start {
		return 0
	}
	return this.prefixProd[len(this.prefixProd)-1] / this.prefixProd[start]
}
