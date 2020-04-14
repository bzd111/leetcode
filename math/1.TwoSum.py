from typing import Dict, List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        data: List[int]
        index: Dict[int, int] = dict()
        for i, v in enumerate(nums):
            if target - v in index.keys():
                data = [index[target - v], i]
            index[v] = i
        return data


if __name__ == "__main__":
    s = Solution()
    print(s.twoSum([2, 2, 3, 4], 4))
    print(s.twoSum([2, 3, 3, 4], 6))
