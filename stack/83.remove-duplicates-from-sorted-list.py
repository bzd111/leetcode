from typing import Optional

#
# @lc app=leetcode id=83 lang=python3
#
# [83] Remove Duplicates from Sorted List
#

# @lc code=start
# Definition for singly-linked list.


class ListNode:
    # def __init__(self, val: int = 0, next: Optional['ListNode'] = None) -> None:
    def __init__(self, val: int = 0, next: Optional["ListNode"] = None) -> None:
        self.val = val
        self.next = next

    def __repr__(self) -> str:
        return f"ListNode(val={self.val}, next={self.next})"

    def __str__(self) -> str:
        res = []
        while self:
            res.append(str(self.val))
            self = self.next  # type: ignore
        return "->".join(res)


class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        result = start = ListNode()
        start.next = head
        while head:
            while head and start.next.val == head.val:
                head = head.next  # type: ignore
            start.next.next = head
            start = start.next
        return result.next  # type: ignore


# @lc code=end
if __name__ == "__main__":
    l1 = ListNode(val=1)
    l2 = ListNode(val=1)
    l3 = ListNode(val=2)
    l4 = ListNode(val=3)
    l5 = ListNode(val=3)
    l1.next = l2
    l2.next = l3
    l3.next = l4
    l4.next = l5
    s = Solution()
    print(s.deleteDuplicates(l1))
