from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) <= 1:
            return 0
        min_price = prices[0]
        max_profit = 0
        for price in prices[1:]:
            min_price = min(price, min_price)
            max_profit = max(max_profit, price - min_price)
        return max_profit


if __name__ == "__main__":
    s = Solution()
    case1 = [7, 1, 5, 3, 6, 4]
    print(s.maxProfit(case1))
    case2 = [7, 6, 4, 3, 1]
    print(s.maxProfit(case2))
    case3 = [2, 1, 4]
    print(s.maxProfit(case3))
