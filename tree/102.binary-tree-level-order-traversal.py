from typeing import List

#
# @lc app=leetcode id=102 lang=python3
#
# [102] Binary Tree Level Order Traversal
#

# @lc code=start
# Definition for a binary tree node.


class TreeNode:
    def __init__(
        self, val: int = 0, left: "TreeNode" = None, right: "TreeNode" = None
    ) -> None:
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return []
        queue = []
        result = []
        queue.append(root)
        while queue:
            size = len(queue)
            values = []
            while size > 0:
                node = queue.pop(0)
                values.append(node.val)
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
                size -= 1
            result.append(values)
        return result


# @lc code=end
