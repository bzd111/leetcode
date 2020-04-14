import string
from typing import List


class Solution:
    def letterCasePermutation(self, S: str) -> List[str]:
        s = S
        ans = []
        length = len(s)

        def dfs(s: List[str], i: int,) -> None:
            if i == length:
                ans.append("".join(s))
                return
            dfs(s, i + 1)
            if s[i] in string.ascii_letters:
                s[i] = chr(ord(s[i]) ^ 1 << 5)
                dfs(s, i + 1)
                s[i] = chr(ord(s[i]) ^ 1 << 5)

        dfs(list(s), 0)
        return ans


if __name__ == "__main__":
    s = Solution()
    t1 = "a1b2"  # ["a1b2", "a1B2", "A1b2", "A1B2"]
    print(s.letterCasePermutation(t1))
    t2 = "3z4"  # ["3z4", "3Z4"]
    print(s.letterCasePermutation(t2))
    t3 = "12345"  # ["12345"]
    print(s.letterCasePermutation(t3))
