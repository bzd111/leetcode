package main

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

// @lc code=start
// *
// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	dfs(root, 0, &res)
	return res
}

func dfs(node *TreeNode, length int, res *[][]int) {
	if node == nil {
		return
	}
	for len(*res) <= length {
		*res = append(*res, []int{})
	}
	(*res)[length] = append((*res)[length], node.Val)
	dfs(node.Left, length+1, res)
	dfs(node.Right, length+1, res)
}

// @lc code=end
