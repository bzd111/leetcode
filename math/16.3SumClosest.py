import sys
from typing import List


class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums = sorted(nums)
        res = 0
        data = sys.maxsize
        for i, num in enumerate(nums):
            l = i + 1
            r = len(nums) - 1
            if i > 0 and nums[i] == nums[i - 1]:
                continue

            while l < r:
                _sum = num + nums[l] + nums[r]
                if target == _sum:
                    return target
                if target < _sum:
                    r -= 1
                    if data > _sum - target:
                        data = _sum - target
                        res = _sum
                else:
                    l += 1
                    if data > target - _sum:
                        data = target - _sum
                        res = _sum
        return res


if __name__ == "__main__":
    s = Solution()

    nums = [-1, 2, 1, -4]
    target = 1
    # target is 2. (-1 + 2 + 1 = 2).
    print(s.threeSumClosest(nums, target))

    nums = [1, 1, 1, 1]
    target = 0
    print(s.threeSumClosest(nums, target))
