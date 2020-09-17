from typing import List

#
# @lc app=leetcode id=221 lang=python3
#
# [221] Maximal Square
#

# @lc code=start


class Solution:
    def maximalSquare(self, matrix: List[List[str]]) -> int:
        if not matrix or len(matrix[0]) == 0:
            return 0
        rows = len(matrix)
        cols = len(matrix[0])
        max_length = 0
        dp = [[0] * (cols + 1) for _ in range(rows + 1)]
        i = 1
        j = 1
        print(rows, cols)
        while i <= rows:
            while j <= cols:
                if matrix[i - 1][j - 1] == "1":
                    dp[i][j] = min(dp[i - 1][j - 1], dp[i - 1][j], dp[i][j - 1]) + 1
                    max_length = max(max_length, dp[i][j])
                j = j + 1
            j = 1
            i = i + 1
        return max_length * max_length


# @lc code=end

if __name__ == "__main__":
    s = Solution()
    matrix = [
        ["1", "0", "1", "0", "0"],
        ["1", "0", "1", "1", "1"],
        ["1", "1", "1", "1", "1"],
        ["1", "0", "0", "1", "0"],
    ]
    print(s.maximalSquare(matrix))  # 4

    matrix = [["1"]]
    print(s.maximalSquare(matrix))  # 1
    matrix = [
        ["1", "1", "1", "1", "1", "1", "1", "1"],
        ["1", "1", "1", "1", "1", "1", "1", "0"],
        ["1", "1", "1", "1", "1", "1", "1", "0"],
        ["1", "1", "1", "1", "1", "0", "0", "0"],
        ["0", "1", "1", "1", "1", "0", "0", "0"],
    ]
    print(s.maximalSquare(matrix))  # 16
