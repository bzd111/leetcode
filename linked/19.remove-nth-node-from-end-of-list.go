/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	fast := head
	prev := dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		prev = prev.Next
		fast = fast.Next
	}
	prev.Next = prev.Next.Next
	return dummy.Next

}
// @lc code=end

