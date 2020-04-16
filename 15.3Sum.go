package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	data := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		fmt.Println("i:", i)
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				data = append(data, []int{nums[i], nums[l], nums[r]})
				// for l+1 < len(nums) && nums[l+1] == nums[l] {
				for l+1 < len(nums) && nums[l+1] == nums[l] {
					l++
				}
				for r-1 >= 0 && nums[r-1] == nums[r] {
					r--
				}
				// for equal
				// l++
				// r--
			}
			if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return data
}

func main() {
	//  Given array nums = [-1, 0, 1, 2, -1, -4]

	// [
	//   [-1, 0, 1],
	//   [-1, -1, 2]
	// ]
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))

	nums = []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
	nums = []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))
}
