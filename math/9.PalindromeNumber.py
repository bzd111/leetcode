class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        s = str(x)
        l = 0
        r = len(s) - 1
        while l < r:
            if s[l] != s[r]:
                return False
            l += 1
            r -= 1
        return True


if __name__ == "__main__":
    s = Solution()
    print(s.isPalindrome(121))  # True
    print(s.isPalindrome(-121))  # False
    print(s.isPalindrome(10))  # False
