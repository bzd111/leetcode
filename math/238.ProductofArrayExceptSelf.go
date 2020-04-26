package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i, _ := range res {
		res[i] = 1
	}
	left := 1
	for i := 0; i < len(nums)-1; i++ {
		left *= nums[i]
		res[i+1] *= left
	}
	right := 1
	for i := len(nums) - 1; i > 0; i-- {
		right *= nums[i]
		res[i-1] *= right
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
