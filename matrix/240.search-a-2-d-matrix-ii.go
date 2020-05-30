/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	r := 0
	c := n - 1
	for r < m && c > 0 {
		if matrix[r][c] == target {
			return True
		}
		if matrix[r][c] > target {
			c = c - 1
		} else {
			r = r + 1
		}
	}
	return false
}

// @lc code=end

