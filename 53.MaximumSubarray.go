package main

import "fmt"

func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := nums[0]
	for _, num := range nums[1:] {
		ans = max(num+ans, num)
		if ans > sum {
			sum = ans
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	m := maxSubArray(nums)
	fmt.Println("m:", m)
}
