from typing import List


class Solution:
    def shiftingLetters(self, S: str, shifts: List[int]) -> str:
        c = 0
        s = list(S)
        l = len(shifts) - 1
        # l = 0
        for i in shifts[::-1]:
            c += i % 26
            s[l] = chr((ord(s[l]) - ord("a") + c) % 26 + ord("a"))
            l -= 1
        return "".join(s)


if __name__ == "__main__":
    # Input: S = "abc", shifts = [3,5,9]
    # Output: "rpl"
    # Explanation:
    # We start with "abc".
    s = Solution()
    print(s.shiftingLetters("abc", [3, 5, 9]))  # 'rpl'
