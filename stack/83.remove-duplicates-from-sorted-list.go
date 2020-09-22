// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func (l *ListNode) String() string {
// 	res := []string{}
// 	for l != nil {
// 		res = append(res, strconv.Itoa(l.Val))
// 		l = l.Next
// 	}
// 	return strings.Join(res, "->")
// }

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// @lc code=end
// func main() {
// 	l1 := &ListNode{Val: 1}
// 	l2 := &ListNode{Val: 1}
// 	l3 := &ListNode{Val: 2}
// 	l4 := &ListNode{Val: 3}
// 	l5 := &ListNode{Val: 3}
// 	l1.Next = l2
// 	l2.Next = l3
// 	l3.Next = l4
// 	l4.Next = l5
// 	fmt.Println(deleteDuplicates(l1))
// }
