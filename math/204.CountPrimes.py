class Solution:
    # https://zh.wikipedia.org/wiki/%E5%9F%83%E6%8B%89%E6%89%98%E6%96%AF%E7%89%B9%E5%B0%BC%E7%AD%9B%E6%B3%95#Python_3.6
    def countPrimes(self, n: int) -> int:
        IsPrime = [True] * n
        for i in range(2, int(n ** 0.5) + 1):
            if IsPrime[i]:
                for j in range(i * i, n, i):
                    IsPrime[j] = False
        return len([x for x in range(2, n) if IsPrime[x]])


if __name__ == "__main__":
    s = Solution()
    n = 10
    print(s.countPrimes(n))
    n = 3
    print(s.countPrimes(n))  # 1
    n = 2
    print(s.countPrimes(n))  # 0
    n = 5
    print(s.countPrimes(n))  # 2
