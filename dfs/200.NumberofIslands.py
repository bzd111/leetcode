from typing import List


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        if len(grid) == 0:
            return 0
        m = len(grid)
        n = len(grid[0])
        ans = 0
        for i, _ in enumerate(range(m)):
            for j, _ in enumerate(range(n)):
                if grid[i][j] == "1":
                    ans += 1
                    self.dfs(grid, m, n, i, j)
        return ans

    def dfs(self, grid: List[List[str]], m: int, n: int, i: int, j: int) -> None:
        if i >= m or i < 0 or j < 0 or j >= n or grid[i][j] == "0":
            return
        grid[i][j] = "0"
        self.dfs(grid, m, n, i + 1, j)
        self.dfs(grid, m, n, i, j + 1)
        self.dfs(grid, m, n, i - 1, j)
        self.dfs(grid, m, n, i, j - 1)


if __name__ == "__main__":
    s = Solution()
    # Input:
    # 11110
    # 11010
    # 11000
    # 00000

    # Output: 1
    grid = [list("11110"), list("11010"), list("11000"), list("00000")]
    print(s.numIslands(grid))

    # Input:
    # 11000
    # 11000
    # 00100
    # 00011

    # Output: 3
    grid = [list("11000"), list("11000"), list("00100"), list("00011")]
    print(s.numIslands(grid))
