from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profit = 0
        for i, price in enumerate(prices[1:], 1):
            profit += max(0, price - prices[i - 1])
        return profit


if __name__ == "__main__":
    s = Solution()
    prices = [7, 1, 5, 3, 6, 4]  # 7
    print(s.maxProfit(prices))
    prices = [1, 2, 3, 4, 5]  # 4
    print(s.maxProfit(prices))
