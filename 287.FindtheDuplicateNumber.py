from typing import List


class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        fast = slow = 0
        slow = nums[slow]
        fast = nums[fast]
        fast = nums[fast]
        while slow != fast:
            slow = nums[slow]
            fast = nums[fast]
            fast = nums[fast]
        slow = 0
        while slow != fast:
            slow = nums[slow]
            fast = nums[fast]
        return slow


if __name__ == "__main__":

    s = Solution()
    l = [1, 3, 4, 2, 2]
    print(s.findDuplicate(l))
    l = [3, 1, 3, 4, 2]
    print(s.findDuplicate(l))
