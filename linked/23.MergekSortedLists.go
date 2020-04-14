package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	res := []string{}
	for l != nil {
		res = append(res, strconv.Itoa(l.Val))
		l = l.Next
	}
	return strings.Join(res, "->")
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)

}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l > r {
		return nil
	}
	if l == r {
		return lists[l]
	}
	if l+1 == r {
		return mergeTwoLists(lists[l], lists[r])
	}
	m := (l + r) / 2
	l1 := merge(lists, l, m)
	l2 := merge(lists, m+1, r)
	return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	tail := &ListNode{}
	dumpy := tail
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return dumpy.Next
}

func main() {
	// 1->4->5,
	// 1->3->4,
	// 2->6
	l11 := &ListNode{Val: 1}
	l12 := &ListNode{Val: 4}
	l13 := &ListNode{Val: 5}
	l11.Next = l12
	l12.Next = l13
	l21 := &ListNode{Val: 1}
	l22 := &ListNode{Val: 3}
	l23 := &ListNode{Val: 4}
	l21.Next = l22
	l22.Next = l23
	l31 := &ListNode{Val: 2}
	l32 := &ListNode{Val: 6}
	l31.Next = l32
	// 1->1->2->3->4->4->5->6
	fmt.Println("l1", l11)
	// fmt.Println("two lists", mergeTwoLists(l11, l21))
	fmt.Println(mergeKLists([]*ListNode{l11, l21, l31}))
}
