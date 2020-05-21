#
# @lc app=leetcode id=20 lang=python3
#
# [20] Valid Parentheses
#
from typing import List


# @lc code=start
class Solution:
    def isValid(self, s: str) -> bool:
        pairs = {")": "(", "]": "[", "}": "{", "(": ")", "[": "]", "{": "}"}
        stack: List[str] = []
        for i in s:
            if not stack:
                stack.append(i)
                continue
            if pairs[i] == stack[-1]:
                stack.pop()
            else:
                stack.append(i)
        if stack:
            return False
        return True


# @lc code=end
