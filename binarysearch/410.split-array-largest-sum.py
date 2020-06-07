#
# @lc app=leetcode id=410 lang=python3
#
# [410] Split Array Largest Sum
#

# @lc code=start
class Solution:
    def splitArray(self, nums: List[int], m: int) -> int:
        l = max(nums)
        r = sum(nums) + 1
        while l < r:
            limit = (r - l) // 2 + l
            if self.min_groups(nums, limit) > m:
                l = limit + 1
            else:
                r = limit
        return l

    def min_groups(self, nums: List[int], limit: int):
        sums = 0
        groups = 1
        for num in nums:
            if sums + num > limit:
                sums = num
                groups += 1
            else:
                sums += num
        return groups


# @lc code=end
