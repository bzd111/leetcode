#
# @lc app=leetcode id=844 lang=python3
#
# [844] Backspace String Compare
#

# @lc code=start
from typing import List


class Solution:
    def backspaceCompare(self, S: str, T: str) -> bool:
        s = self.realyString(S)
        t = self.realyString(T)
        return s == t

    def realyString(self, originString: str) -> List[str]:
        result: List[str] = []
        for i in originString:
            if i == "#":
                if len(result) > 0:
                    result.pop()
            else:
                result.append(i)
        return result


# @lc code=end
