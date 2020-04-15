class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        l1 = len(word1)
        l2 = len(word2)
        self.d = [[0] * (l2 + 1) for _ in range(l1 + 1)]
        # self.d = [[-1 for _ in range(l2 + 1)] for _ in range(l1 + 1)]

        return self.dfs(word1, word2, l1, l2)

    def dfs(self, word1, word2, l1, l2):
        if not l1:
            return l2
        if not l2:
            return l1
        if self.d[l1][l2] > 0:
            return self.d[l1][l2]
        if word1[l1 - 1] == word2[l2 - 1]:
            ans = self.dfs(word1, word2, l1 - 1, l2 - 1)
        else:
            ans = (
                min(
                    min(
                        self.dfs(word1, word2, l1 - 1, l2 - 1),
                        self.dfs(word1, word2, l1 - 1, l2),
                    ),
                    self.dfs(word1, word2, l1, l2 - 1),
                )
                + 1
            )
        self.d[l1][l2] = ans
        return ans


if __name__ == "__main__":
    # Input: word1 = "horse", word2 = "ros"
    # Output: 3
    # Explanation:
    # horse -> rorse (replace 'h' with 'r')
    # rorse -> rose (remove 'r')
    # rose -> ros (remove 'e')
    # Input: word1 = "intention", word2 = "execution"
    # Output: 5
    # Explanation:
    # intention -> inention (remove 't')
    # inention -> enention (replace 'i' with 'e')
    # enention -> exention (replace 'n' with 'x')
    # exention -> exection (replace 'n' with 'c')
    # exection -> execution (insert 'u')
    s = Solution()
    word1 = "horse"
    word2 = "ros"
    print(s.minDistance(word1, word2))
