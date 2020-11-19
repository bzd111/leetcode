package main

import "math"

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */

// @lc code=start
// *
// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isValidBST(root *TreeNode) bool {
	return isValidbst(root, math.Inf(-1), math.Inf(1))
}

func isValidbst(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v < max && v > min && isValidbst(root.Left, min, v) && isValidbst(root.Right, v, max)
}

// @lc code=end

// class Solution {
//   public boolean isValidBST(TreeNode root) {
//     return isValidBST(root, null, null);
//   }

//   private boolean isValidBST(TreeNode root, Integer min, Integer max) {
//     if (root == null) return true;
//     if ((min != null && root.val <= min)
//       ||(max != null && root.val >= max))
//         return false;

//     return isValidBST(root.left, min, root.val)
//         && isValidBST(root.right, root.val, max);
//   }
// }
