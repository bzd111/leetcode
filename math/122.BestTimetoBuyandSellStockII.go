package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	for i, price := range prices[1:] {
		profit += max(0, price-prices[i])
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4} //  7
	fmt.Println(maxProfit(prices))
	prices = []int{1, 2, 3, 4, 5} // 4
	fmt.Println(maxProfit(prices))
}
