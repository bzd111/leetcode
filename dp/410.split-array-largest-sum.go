package main

import "fmt"

/*
 * @lc app=leetcode id=410 lang=golang
 *
 * [410] Split Array Largest Sum
 */

// @lc code=start
func splitArray(nums []int, m int) int {
	n := len(nums)
	sums := make([]int, n)
	sums[0] = nums[0]
	dp := make([][]int, m+1)
	// var max_int int64 = 2147483648
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 2147483648
		}
	}
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	for i := 0; i < n; i++ {
		dp[1][i] = sums[i]
	}

	for i := 2; i <= m; i++ {
		for j := i - 1; j < n; j++ {
			for k := 0; k < j; k++ {
				dp[i][j] = min(dp[i][j], max(dp[i-1][k], sums[j]-sums[k]))
			}
		}
	}
	return dp[m][n-1]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// func main() {
// 	nums := []int{7, 2, 5, 10, 8}
// 	m := 2
// 	fmt.Println(splitArray(nums, m))
// }

// @lc code=end
