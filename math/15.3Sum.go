package main

import (
	"fmt"
	"sort"
)

// func threeSum(nums []int) [][]int {
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := range nums {
		j, k := i+1, len(nums)-1
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j, k = move(nums, j, k)
			}
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			}
		}
	}
	return res
}

func move(nums []int, j, k int) (int, int) {
	for j < k {
		switch {
		case nums[j] == nums[j+1]:
			j++
		case nums[k] == nums[k-1]:
			k--
		default:
			j++
			k--
			return j, k

		}
	}
	return j, k
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	nums = []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}

// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]
