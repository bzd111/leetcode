package main

/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 */

// @lc code=start
// *
// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	m := l + (r-l)/2
	root := TreeNode{Val: nums[m]}
	root.Left = buildBST(nums, l, m-1)
	root.Right = buildBST(nums, m+1, r)
	return &root
}

// @lc code=end
