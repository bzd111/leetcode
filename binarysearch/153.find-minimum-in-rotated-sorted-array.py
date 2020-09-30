from typing import List

#
# @lc app=leetcode id=153 lang=python3
#
# [153] Find Minimum in Rotated Sorted Array
#

# @lc code=start


class Solution:
    def findMin(self, nums: List[int]) -> int:
        lo = 0
        hi = len(nums) - 1
        if nums[lo] < nums[hi]:
            return nums[lo]
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if mid > 0 and nums[mid] < nums[mid - 1]:
                return nums[mid]
            if nums[mid] > nums[mid + 1]:
                return nums[mid + 1]
            if nums[lo] < nums[mid]:
                lo = mid + 1
            else:
                hi = mid - 1

        return nums[lo]


# @lc code=end
if __name__ == "__main__":
    nums = [3, 4, 5, 1, 2]
    s = Solution()
    print(s.findMin(nums))
