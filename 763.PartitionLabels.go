package main

import "fmt"

func partitionLabels(S string) []int {
	s := S
	result := []int{}
	data := make([]int, 128)
	for i, v := range s {
		data[int(v)] = i
	}
	start := 0
	end := 0
	for i, v := range s {
		end = max(end, data[v])
		if end == i {
			result = append(result, end-start+1)
			start = end + 1
			end = 0
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Input: S = "ababcbacadefegdehijhklij"
	// Output: [9,7,8]
	// Explanation:
	// The partition is "ababcbaca", "defegde", "hijhklij".
	t := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(t))
}
