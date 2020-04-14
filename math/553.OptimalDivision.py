from typing import List


class Solution:
    def optimalDivision(self, nums: List[int]) -> str:
        if len(nums) == 1:
            return str(nums[0])
        if len(nums) == 2:
            return str(nums[0]) + "/" + str(nums[1])
        nums = list(map(str, nums))
        a = "/".join(nums[1:])
        return nums[0] + "/(" + a + ")"


if __name__ == "__main__":
    # Input: [1000,100,10,2]
    # Output: "1000/(100/10/2)"
    s = Solution()
    nums = [1000, 100, 10, 2]
    print(s.optimalDivision(nums))
    nums = [2]
    print(s.optimalDivision(nums))
    nums = [10, 2]
    print(s.optimalDivision(nums))
