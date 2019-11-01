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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	begin := &ListNode{}
	list := begin
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			list.Next = l1
			l1 = l1.Next
		} else {
			list.Next = l2
			l2 = l2.Next
		}
		list = list.Next

	}
	if l1 != nil {
		list.Next = l1
	}
	if l2 != nil {
		list.Next = l2
	}

	return begin.Next
}

func main() {
	// 1->2->4, 1->3->4
	// 1->1->2->3->4->4
	l11 := &ListNode{Val: 1}
	l12 := &ListNode{Val: 2}
	l13 := &ListNode{Val: 4}
	l11.Next = l12
	l12.Next = l13
	l21 := &ListNode{Val: 1}
	l22 := &ListNode{Val: 3}
	l23 := &ListNode{Val: 4}
	l21.Next = l22
	l22.Next = l23
	fmt.Println(mergeTwoLists(l11, l21))
}
