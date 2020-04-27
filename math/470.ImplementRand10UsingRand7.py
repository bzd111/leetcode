import sys


def rand7() -> int:
    ...


class Solution:
    def rand10(self) -> int:
        index = sys.maxsize
        while index > 40:
            index = 7 * (rand7() - 1) + rand7() - 1
        return index % 10 + 1
