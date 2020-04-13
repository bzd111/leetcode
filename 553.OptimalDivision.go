package main

import (
	"fmt"
	"strconv"
)

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}
	r := []string{}
	for _, v := range nums {
		r = append(r, strconv.Itoa(v))
	}
	result := r[0] + "/(" + r[1]
	for i := 2; i < len(r); i++ {
		result += "/" + r[i]
	}
	result = result + ")"
	return result
}

func main() {

	nums := []int{1000, 100, 10, 2}
	fmt.Println(optimalDivision(nums))
	nums = []int{10, 2}
	fmt.Println(optimalDivision(nums))
	nums = []int{2}
	fmt.Println(optimalDivision(nums))
}
