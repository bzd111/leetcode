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

func reverseList(head *ListNode) *ListNode {
	// var prev  &ListNode = nil
	var prev *ListNode
	for head != nil {
		_next := head.Next
		head.Next = prev
		prev = head
		head = _next
	}
	return prev
}

func main() {
	l5 := &ListNode{
		Val: 5,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l5,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	fmt.Println(reverseList(l1))
}
