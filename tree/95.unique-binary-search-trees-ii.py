from typing import List

#
# @lc app=leetcode id=95 lang=python3
#
# [95] Unique Binary Search Trees II
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def generateTrees(self, n: int) -> List[TreeNode]:
        if n == 0:
            return []

        return self.helper(1, n)

    def helper(self, l: int, r: int) -> List[TreeNode]:
        result: List[TreeNode] = []
        if l > r:
            result.append(None)
            return result
        for i in range(l, r + 1):
            for left in self.helper(l, i - 1):
                for right in self.helper(i + 1, r):
                    root = TreeNode(val=i, left=left, right=right)
                    result.append(root)
        return result


# @lc code=end
