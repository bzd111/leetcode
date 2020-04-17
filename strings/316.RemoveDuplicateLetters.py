from typing import List


class Solution:
    def removeDuplicateLetters(self, s: str) -> str:
        # c = Counter(s)
        c = [0] * 26
        use = [0] * 26
        res: List[str] = []
        for i in s:
            c[ord(i) - ord("a")] += 1

        for i in s:
            c[ord(i) - ord("a")] -= 1

            if use[ord(i) - ord("a")]:
                continue

            while len(res) > 0 and res[-1] > i and c[ord(res[-1]) - ord("a")] > 0:
                use[ord(res[-1]) - ord("a")] = False
                res.remove(res[-1])
            res.append(i)
            use[ord(i) - ord("a")] = True

        return "".join(res)


if __name__ == "__main__":
    # Input: "bcabc"
    # Output: "abc"

    # Input: "cbacdcbc"
    # Output: "acdb"
    s = Solution()
    print(s.removeDuplicateLetters("bcabc"))
    print(s.removeDuplicateLetters("cbacdcbc"))

    print(s.removeDuplicateLetters("cbbbcaa"))  # "bca"
