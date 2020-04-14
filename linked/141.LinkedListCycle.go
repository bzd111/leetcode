package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	rabbit := head
	turtle := head
	for rabbit != nil && rabbit.Next != nil {
		rabbit = rabbit.Next.Next
		turtle = turtle.Next
		if rabbit == turtle {
			return true
		}
	}
	return false
}

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l2
	fmt.Println(hasCycle(l1))

	l1 = nil
	fmt.Println(hasCycle(l1))
}
