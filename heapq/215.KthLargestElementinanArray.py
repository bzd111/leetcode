import heapq
from typing import List


class Solution:
    def __init__(self) -> None:
        self.minheap: List[int] = []

    def findKthLargest(self, nums: List[int], k: int) -> int:
        for i in nums:
            self.push(i)
        for _ in range(len(nums) - k):
            heapq.heappop(self.minheap)
        return heapq.heappop(self.minheap)

    def push(self, val: int) -> None:
        heapq.heappush(self.minheap, val)


if __name__ == "__main__":
    s = Solution()
    nums = [3, 2, 1, 5, 6, 4]
    k = 2
    print(s.findKthLargest(nums, k))  # 5
    s = Solution()
    nums = [3, 2, 3, 1, 2, 4, 5, 5, 6]
    k = 4
    print(s.findKthLargest(nums, k))  # 4
