package main

import (
	"container/heap"
	"fmt"
)

type MaxHeaq []int

func findKthLargest(nums []int, k int) int {
	temp := MaxHeaq(nums)
	h := &temp
	heap.Init(h)
	for i := 1; i < k; i++ {
		heap.Remove(h, 0)
	}
	return (*h)[0]
}

func (h MaxHeaq) Len() int {
	return len(h)
}

func (h *MaxHeaq) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeaq) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}

func (h MaxHeaq) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeaq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k)) // 5
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(findKthLargest(nums, k)) //4
}
