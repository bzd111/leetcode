from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums = sorted(nums)
        data = []
        for i in range(len(nums)):
            l = i + 1
            r = len(nums) - 1
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            while l < r:
                _sum = nums[i] + nums[l] + nums[r]
                if _sum == 0:
                    data.append([nums[i], nums[l], nums[r]])
                    while l + 1 < len(nums) and nums[l] == nums[l + 1]:
                        l += 1
                    while r - 1 >= 0 and nums[r] == nums[r - 1]:
                        r -= 1
                if _sum > 0:
                    r -= 1
                else:
                    l += 1
        return data


if __name__ == "__main__":
    s = Solution()

    nums = [-1, 0, 1, 2, -1, -4]
    print(s.threeSum(nums))
    # [[-1,-1,2],[-1,0,1]]
    nums = [0, 0, 0, 0]
    # [[0,0,0]]
    print(s.threeSum(nums))
    nums = [-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6]
    # [[-4,-2,6],[-4,0,4],[-4,1,3],[-4,2,2],[-2,-2,4],[-2,0,2]]
    print(s.threeSum(nums))
