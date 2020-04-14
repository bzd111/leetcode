from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        ans = nums[0]
        result = nums[0]
        for num in nums[1:]:
            result = max(result + num, num)
            if result > ans:
                ans = result
        return ans


if __name__ == "__main__":
    s = Solution()
    case1 = [-2, 1, -3, 4, -1, 2, 1, -5, 4]  # 6
    print(s.maxSubArray(case1))
