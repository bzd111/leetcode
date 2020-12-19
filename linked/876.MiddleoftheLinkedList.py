# Definition for singly-linked list.


class Solution:
    def middleNode(self, head: ListNode) -> ListNode:
        fast = head
        low = head
        while fast and fast.next:
            fast = fast.next.next
            low = low.next
        if fast:
            low = low.next
        return low
