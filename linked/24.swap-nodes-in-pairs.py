from typing import List, Optional

#
# @lc app=leetcode id=24 lang=python3
#
# [24] Swap Nodes in Pairs
#

# @lc code=start
# Definition for singly-linked list.


class ListNode:
    def __init__(
        self, val: int = 0, next: Optional[ListNode] = None  # noqa: F821
    ) -> None:
        self.val = val
        self.next = next

    def __str__(self) -> str:
        res: List[str] = []
        res.append(str(self.val))
        while self.next:
            res.append(str(self.next.val))
            self.next = self.next.next
        return "->".join(res)


class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        if not head or not head.next:
            return head
        dummy = ListNode(0)
        dummy.next = head
        head = dummy
        while head.next and head.next.next:
            n1, n2 = head.next, head.next.next
            n1.next = n2.next
            n2.next = n1
            head.next = n2
            head = n1
        return dummy.next


# @lc code=end
if __name__ == "__main__":
    l11 = ListNode(1)
    l12 = ListNode(4)
    l13 = ListNode(5)
    l14 = ListNode(6)
    l15 = ListNode(8)
    l11.next = l12
    l12.next = l13
    l13.next = l14
    l14.next = l15
    s = Solution()
    print(s.swapPairs(l11))
