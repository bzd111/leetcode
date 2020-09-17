package main

/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows := len(matrix)
	cols := len(matrix[0])
	length := 0
	dp := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		dp[i] = make([]int, cols+1)
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				length = max(dp[i][j], length)
			}
		}
	}
	return length * length
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
// func main() {
// 	matrix := [][]byte{[]byte{'0', '1', '1', '0', '0'}, []byte{'0', '1', '1', '0', '0'}}
// 	fmt.Println(maximalSquare(matrix))
// }
