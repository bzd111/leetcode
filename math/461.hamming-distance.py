#
# @lc app=leetcode id=461 lang=python3
#
# [461] Hamming Distance
#

# @lc code=start
class Solution:
    def hammingDistance(self, x: int, y: int) -> int:
        res = 0
        a = x ^ y
        while a:
            if a % 2 == 1:
                res += 1
            a = a >> 1
        return res


# @lc code=end

if __name__ == '__main__':
    s = Solution()
    print(s.hammingDistance(1, 4))
