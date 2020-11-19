#
# @lc app=leetcode id=98 lang=python3
#
# [98] Validate Binary Search Tree
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None) -> None:  # type: ignore
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        self.prev: TreeNode = None  # type: ignore
        return self.inOrder(root)

    def inOrder(self, root: TreeNode) -> bool:
        if not root:
            return True
        if not self.inOrder(root.left):
            return False
        if self.prev and root.val <= self.prev.val:
            return False
        self.prev = root
        return self.inOrder(root.right)


# @lc code=end
