from typing import List


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        size = len(nums)
        output = [1] * size
        left = 1
        for x in range(size - 1):
            left *= nums[x]
            output[x + 1] *= left
        right = 1
        for x in range(size - 1, 0, -1):
            right *= nums[x]
            output[x - 1] *= right
        return output


if __name__ == "__main__":
    # Input:  [1,2,3,4]
    # Output: [24,12,8,6]
    nums = [1, 2, 3, 4]
    s = Solution()
    print(s.productExceptSelf(nums))
