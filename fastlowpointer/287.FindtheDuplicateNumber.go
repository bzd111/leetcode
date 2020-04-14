package main

import "fmt"

func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	slow = nums[slow]
	fast = nums[fast]
	fast = nums[fast]
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
		fast = nums[fast]
	}
	anther := 0
	for anther != slow {
		anther = nums[anther]
		slow = nums[slow]
	}
	return slow
}

func main() {
	nums := []int{1, 3, 4, 2, 2} // 2
	fmt.Println(findDuplicate(nums))
}
