from typing import List


class Solution:
    def numTeams(self, rating: List[int]) -> int:
        ans = 0
        for i in range(0, len(rating)):
            for j in range(i + 1, len(rating)):
                for k in range(j + 1, len(rating)):
                    if (
                        rating[i] > rating[j]
                        and rating[j] > rating[k]
                        or rating[i] < rating[j]
                        and rating[j] < rating[k]
                    ):
                        ans += 1
        return ans


if __name__ == "__main__":
    s = Solution()
    a = [2, 5, 3, 4, 1]  # 3
    print(s.numTeams(a))
    a = [
        2,
        1,
        3,
    ]  # 0
    print(s.numTeams(a))
    a = [1, 2, 3, 4]  # 4
    print(s.numTeams(a))
