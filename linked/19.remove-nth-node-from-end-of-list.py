#
# @lc app=leetcode id=19 lang=python3
#
# [19] Remove Nth Node From End of List
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:  # type:ignore
        dummy = ListNode() # type:ignore
        dummy.next = head
        fast, prev = head, dummy
        for _ in range(n):
            fast = fast.next
        while fast:
            fast = fast.next
            prev = prev.next
        prev.next = prev.next.next
        return dummy.next


# @lc code=end
