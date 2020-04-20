package main

import (
	"fmt"
)

func judgePoint24(nums []int) bool {
	l := []float32{}
	for _, num := range nums {
		l = append(l, float32(num))
	}
	return calclate(l)
}

func calclate(nums []float32) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		diff := nums[0] - 24
		if diff < 0 {
			diff = -diff
		}
		if diff < 0.0001 {
			return true
		}
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			b := []float32{}
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					b = append(b, nums[k])
				}
			}
			if calclate(append(b, nums[i]+nums[j])) {
				return true
			}
			if calclate(append(b, nums[i]-nums[j])) {
				return true
			}
			if calclate(append(b, nums[i]*nums[j])) {
				return true
			}
			if nums[j] != 0 {
				if calclate(append(b, nums[i]/nums[j])) {
					return true
				}

			}
		}
	}
	return false
}

func main() {
	nums := []int{4, 1, 8, 7}
	fmt.Println(judgePoint24(nums))

	nums = []int{1, 2, 1, 2}
	fmt.Println(judgePoint24(nums))

	nums = []int{3, 4, 6, 7}
	fmt.Println(judgePoint24(nums))
}
