package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for k, _ := range isPrime {
		isPrime[k] = true
	}
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j = j + i {
				isPrime[j] = false
			}
		}
	}
	ans := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			ans += 1
		}
	}
	return ans
}

func main() {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(2))
	fmt.Println(countPrimes(3))
	fmt.Println(countPrimes(5))
}
