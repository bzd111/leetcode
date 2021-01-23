from typing import List, Optional

#
# @lc app=leetcode id=108 lang=python3
#
# [108] Convert Sorted Array to Binary Search Tree
#

# @lc code=start
# Definition for a binary tree node.


class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:  # type: ignore
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        return self.buildBST(0, len(nums) - 1, nums)

    def buildBST(self, l: int, r: int, nums: List[int]) -> Optional[TreeNode]:
        if l > r:
            return None
        m = l + (r - l) // 2
        root = TreeNode(val=nums[m])
        root.left = self.buildBST(l, m - 1, nums)
        root.right = self.buildBST(m + 1, r, nums)
        return root


# @lc code=end
