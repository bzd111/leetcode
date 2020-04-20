from typing import Any, List


class Solution:
    def judgePoint24(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return abs(nums[0] - 24) < 0.00001
        for i, x in enumerate(nums):
            for j, y in enumerate(nums[i + 1 :], i + 1):
                b = []
                for k, z in enumerate(nums):
                    if k != i and k != j:
                        b.append(z)
                for j in self.computed(float(x), float(y)):
                    if self.judgePoint24([j, *b]):
                        return True
        return False

    def computed(self, x: float, y: float) -> List[Any]:
        result = [x - y, y - x, x * y, x + y]
        if x == 0 or y == 0:
            result.append(0)
        else:
            result.extend([x / y, y / x])
        return result


if __name__ == "__main__":
    s = Solution()
    # Input: [4, 1, 8, 7]
    # Output: True
    # Explanation: (8-4) * (7-1) = 24
    i = [4, 1, 8, 7]
    print(s.judgePoint24(i))

    # Input: [1, 2, 1, 2]
    # Output: False
    i = [1, 2, 1, 2]
    print(s.judgePoint24(i))
