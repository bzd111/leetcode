package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	slow = reverseLinked(slow)
	for slow != nil {
		if slow.Val != head.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}
	return true
}

func reverseLinked(head *ListNode) *ListNode {
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
	l5 := &ListNode{Val: 1}
	l4 := &ListNode{Val: 2, Next: l5}
	l3 := &ListNode{Val: 3, Next: l4}
	l2 := &ListNode{Val: 2, Next: l3}
	l1 := &ListNode{Val: 1, Next: l2}
	fmt.Println(isPalindrome(l1))
}
