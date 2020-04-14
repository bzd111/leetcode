# Definition for singly-linked list.
class ListNode:
    def __init__(self, x: int) -> None:
        self.val = x
        self.next = None


class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        rabbit = head
        turtle = head
        while rabbit is not None and rabbit.next is not None:
            rabbit = rabbit.next.next
            turtle = turtle.next
            if rabbit == turtle:
                return True
        return False


if __name__ == "__main__":
    # TestCase 1
    l1 = ListNode(3)
    l2 = ListNode(2)
    l3 = ListNode(0)
    l4 = ListNode(4)
    l1.next = l2  # type: ignore
    l2.next = l3  # type: ignore
    l3.next = l4  # type: ignore
    l4.next = l2  # type: ignore
    s = Solution()
    print(s.hasCycle(l1))

    # TestCase 2
    l1 = ListNode(1)
    l2 = ListNode(2)
    print(s.hasCycle(l1))
