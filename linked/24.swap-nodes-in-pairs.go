package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
// *
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy
	for head.Next != nil && head.Next.Next != nil {
		n1, n2 := head.Next, head.Next.Next
		n1.Next = n2.Next
		n2.Next = n1
		head.Next = n2
		head = n1
	}
	return dummy.Next
}

func (l *ListNode) String() string {
	res := []string{}
	for l != nil {
		res = append(res, strconv.Itoa(l.Val))
		l = l.Next
	}
	return strings.Join(res, "->")
}

// @lc code=end

func main() {
	l11 := &ListNode{Val: 1}
	l12 := &ListNode{Val: 4}
	l13 := &ListNode{Val: 5}
	l14 := &ListNode{Val: 8}
	l11.Next = l12
	l12.Next = l13
	l13.Next = l14
	fmt.Println(swapPairs(l11))
}
