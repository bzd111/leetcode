from typing import List, Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val: int = 0, next: "ListNode" = None) -> None:
        self.val = val
        self.next = next

    def __str__(self) -> str:
        res: List[str] = []
        while self:
            res.append(str(self.val))
            self = self.next  # type: ignore
        return "->".join(res)


class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        fast: Optional[ListNode] = head
        slow: Optional[ListNode] = head
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next  # type:ignore
        if fast:
            slow = slow.next  # type:ignore
        slow = reverse_linked(slow)
        while slow:
            if slow.val != head.val:
                return False
            slow = slow.next
            head = head.next  # type:ignore
        return True


def reverse_linked(head: Optional[ListNode]) -> Optional[ListNode]:
    prev: Optional[ListNode] = None
    while head:
        _next: Optional[ListNode] = head.next
        head.next = prev
        prev = head
        head = _next
    return prev


if __name__ == "__main__":
    s = Solution()
    l1 = ListNode(val=1)
    l2 = ListNode(val=2)
    l3 = ListNode(val=3)
    l4 = ListNode(val=3)
    l5 = ListNode(val=2)
    l6 = ListNode(val=1)
    l1.next = l2
    l2.next = l3
    l3.next = l4
    l4.next = l5
    l5.next = l6
    print(s.isPalindrome(l1))
    # print(l1)

    l1 = ListNode(val=1)
    l2 = ListNode(val=2)
    l3 = ListNode(val=1)
    l1.next = l2
    l2.next = l3
    print(s.isPalindrome(l1))
