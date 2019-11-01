from typing import List, Union


class ListNode:
    def __str__(self) -> str:
        res: List[str] = []
        res.append(str(self.val))
        while self.next:  # type: ignore
            res.append(str(self.next.val))  # type: ignore
            self.next = self.next.next  # type: ignore
        return "->".join(res)

    def __init__(self, x: int) -> None:
        self.val = x
        # self.next: Union[ListNode] = None
        self.next: Union["ListNode"] = None  # type: ignore


class Solution(object):
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        self.b = start = ListNode(0)
        while l1 and l2:
            if l1.val < l2.val:
                start.next = l1
                l1 = l1.next
            else:
                start.next = l2
                l2 = l2.next
            start = start.next
        if l1:
            start.next = l1
        if l2:
            start.next = l2
        return self.b.next


if __name__ == "__main__":
    # 1->2->4, 1->3->4
    # 1->1->2->3->4->4
    l1 = ListNode(1)
    l12 = ListNode(2)
    l1.next = l12  # type: ignore
    l13 = ListNode(4)
    l12.next = l13  # type: ignore

    l2 = ListNode(1)
    l22 = ListNode(3)
    l2.next = l22  # type: ignore
    l23 = ListNode(4)
    l22.next = l23  # type: ignore
    s = Solution()
    res = s.mergeTwoLists(l1, l2)
    print(res)
