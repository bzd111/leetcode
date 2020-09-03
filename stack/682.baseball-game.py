from typing import List

#
# @lc app=leetcode id=682 lang=python3
#
# [682] Baseball Game
#

# @lc code=start


class Solution:
    def calPoints(self, ops: List[str]) -> int:
        result: List[int] = []
        for s in ops:
            if s == "C":
                result.pop()
            elif s == "+":
                s1 = result.pop()
                s2 = result.pop()
                result.append(s2)
                result.append(s1)
                result.append(s1 + s2)
            elif s == "D":
                s1 = result.pop()
                result.append(s1)
                result.append(s1 * 2)
            else:
                result.append(int(s))
        return sum(result)


# @lc code=end

if __name__ == "__main__":
    s = Solution()
    ops = ["5", "-2", "4", "C", "D", "9", "+", "+"]

    print(s.calPoints(ops))
