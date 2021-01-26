/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 */

// @lc code=start
// *
// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return helper(1, n)
}

func helper(l, r int) []*TreeNode {
	result := []*TreeNode{}
	if l > r {
		result = append(result, nil)
		return result
	}
	for i := l; i <= r; i++ {
		for _, left := range helper(l, i-1) {
			for _, right := range helper(i+1, r) {
				root := TreeNode{Val: i}
				root.Left = left
				root.Right = right
				result = append(result, &root)
			}
		}
	}
	return result
}

// @lc code=end
