#
# @lc app=leetcode id=394 lang=python3
#
# [394] Decode String
#

# @lc code=start


class Solution:
    def decodeString(self, s: str) -> str:
        numStack = []
        strStack = []
        tail = ""
        i = 0
        while i < len(s):
            if s[i].isdigit():
                n = int(s[i])
                while s[i + 1].isdigit():
                    n = 10 * n + int(s[i + 1])
                    i += 1
                numStack.append(n)
            elif s[i] == "[":
                strStack.append(tail)
                tail = ""
            elif s[i] == "]":
                tmp = strStack.pop()
                times = numStack.pop()
                tmp = tmp + tail * times
                # strStack.append(tmp)
                tail = tmp
            else:
                tail += s[i]
            i += 1
        return tail


# @lc code=end
if __name__ == "__main__":
    s = Solution()
    print(s.decodeString("3[a]2[bc]"))
