package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 2, 9, 11}, 4))
	fmt.Println(twoSum([]int{2, 1, 9, 11}, 3))
}

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := index[target-v]; ok {
			return []int{j, i}
		}
		index[v] = i
	}
	return nil
}
