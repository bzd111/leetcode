from typing import List


class Solution:
    def partitionLabels(self, S: str) -> List[int]:
        s = S
        data = {}
        for i, v in enumerate(s):
            data[v] = i
        start = 0
        end = 0
        result: List[int] = []
        for i, v in enumerate(s):
            end = max(end, data[v])
            print("start", start)
            if end == i:
                result.append(end - start + 1)
                start = end + 1
        return result


if __name__ == "__main__":
    # Input: S = "ababcbacadefegdehijhklij"
    # Output: [9,7,8]
    # Explanation:
    # The partition is "ababcbaca", "defegde", "hijhklij".
    t = "ababcbacadefegdehijhklij"
    s = Solution()
    print(s.partitionLabels(t))
