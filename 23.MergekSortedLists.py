import heapq
from typing import List, Tuple


class ListNode:
    def __str__(self) -> str:
        res = []
        res.append(str(self.val))
        while self.next:  # type:ignore
            res.append(str(self.next.val))  # type:ignore
            self.next = self.next.next  # type:ignore
        return "->".join(res)

    def __init__(self, x: int) -> None:
        self.val = x
        self.next = None

    # def __lt__(self, other):
    #     return self.val < other.val


class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        ListNode.__lt__ = lambda self, other: self.val < other.val  # type:ignore
        data: List[Tuple[int, ListNode]] = []
        dumpy = ListNode(0)
        begin = dumpy
        heapq.heapify(data)
        for l in lists:
            if l:
                heapq.heappush(data, (l.val, l))
        while data:
            val, head = heapq.heappop(data)
            dumpy.next = head
            dumpy = dumpy.next
            if head.next:
                heapq.heappush(data, (head.next.val, head.next))

        return begin.next


if __name__ == "__main__":
    #   1->4->5,
    # 1->3->4,
    # 2->6
    l11 = ListNode(1)
    l12 = ListNode(4)
    l13 = ListNode(5)
    l21 = ListNode(1)
    l22 = ListNode(3)
    l23 = ListNode(4)
    l31 = ListNode(2)
    l32 = ListNode(6)
    l11.next = l12  # type: ignore
    l12.next = l13  # type: ignore
    l21.next = l22  # type: ignore
    l22.next = l23  # type: ignore
    l31.next = l32  # type: ignore
    s = Solution()
    print(s.mergeKLists([l11, l21, l31]))
    # 1->1->2->3->4->4->5->6
