class ListNode:
    def __init__(self, x: int) -> None:
        self.val = x
        self.next = None

    def __str__(self: "ListNode") -> str:
        res = []
        while self:
            res.append(self.val)
            self = self.next  # type: ignore
        s = map(str, res)
        return "->".join(s)

    # def __str__(self):
    # result = []
    # while self:
    #     result.append(str(self.val))
    #     self = self.next
    # return '->'.join(result)


class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        prev = None
        while head:
            _next = head.next
            head.next = prev
            prev = head
            head = _next  # type: ignore
        return prev  # type: ignore


if __name__ == "__main__":

    # Input: 1->2->3->4->5->NULL
    # Output: 5->4->3->2->1->NULL
    l1 = ListNode(1)
    l2 = ListNode(2)
    l3 = ListNode(3)
    l4 = ListNode(4)
    l5 = ListNode(5)
    l1.next = l2  # type: ignore
    l2.next = l3  # type: ignore
    l3.next = l4  # type: ignore
    l4.next = l5  # type: ignore

    print(l1)
    s = Solution()
    print(s.reverseList(l1))
