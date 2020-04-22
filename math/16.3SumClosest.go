package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, data := 0, math.MaxInt64

	for i := range nums {
		j, k := i+1, len(nums)-1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if sum > target {
				k--
				if data > sum-target {
					data = sum - target
					res = sum
				}
			} else if sum < target {
				j++
				if data > target-sum {
					data = target - sum
					res = sum
				}
			}
		}
	}
	return res
}
