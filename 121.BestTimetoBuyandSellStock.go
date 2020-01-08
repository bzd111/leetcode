package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices[1:] {
		minPrice = min(price, minPrice)
		maxProfit = max(maxProfit, price-minPrice)
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4} // 5
	fmt.Println(maxProfit(prices))
}
