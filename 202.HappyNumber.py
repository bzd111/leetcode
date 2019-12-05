class Solution:
    def isHappy(self, n: int) -> bool:
        slow = fast = n
        slow = calc(slow)
        fast = calc(fast)
        fast = calc(fast)
        while slow != fast:
            slow = calc(slow)
            fast = calc(fast)
            fast = calc(fast)
        if slow == 1:
            return True
        return False


def calc(n: int) -> int:
    sum = 0
    while n > 0:
        l = n % 10
        sum += l * l
        n //= 10
    return sum


if __name__ == "__main__":
    s = Solution()
    print(s.isHappy(19))
