package main

import "fmt"

func singleNumber(nums []int) int {
	num := nums[0]
	for _, value := range nums[1:] {
		num = num ^ value
	}
	return num
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
