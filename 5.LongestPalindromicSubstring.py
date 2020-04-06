class Solution:
    def longestPalindrome(self, s: str) -> str:
        length = 0
        start = 0
        for i in range(len(s)):
            cur = max(self.get_lens(s, i, i), self.get_lens(s, i, i + 1))
            if cur > length:
                length = cur
                start = i - (cur - 1) // 2
        return s[start : length + start]

    def get_lens(self, s: str, l: int, r: int) -> int:
        while l >= 0 and r < len(s) and s[l] == s[r]:
            l -= 1
            r += 1
        return r - l - 1


if __name__ == "__main__":
    s = Solution()
    t = "babad"
    print(s.longestPalindrome(t))
    t = "cbbd"
    print(s.longestPalindrome(t))
    t = "a"
    print(s.longestPalindrome(t))
    t = "ccc"
    print(s.longestPalindrome(t))
