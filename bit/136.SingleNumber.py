from typing import List


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        single = nums[0]
        for num in nums[1:]:
            single ^= num
        return single


if __name__ == "__main__":
    s = Solution()
    nums = [4, 1, 2, 1, 2]  # 4
    print(s.singleNumber(nums))
