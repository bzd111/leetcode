package main

import (
	"fmt"
	"math"
)

func myPow(x float64, n int) float64 {
	positive := true
	if n < 0 {
		positive = false
	}
	total := math.Pow(x, math.Abs(float64(n)))
	if positive {
		return total
	} else {
		return 1 / total
	}

}

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}
