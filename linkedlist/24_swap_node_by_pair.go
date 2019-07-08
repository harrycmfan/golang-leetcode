package linkedlist

/*
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.



Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{}
	curr := dummyHead
	groupStart := head
	curr.Next = groupStart.Next
	groupEnd := groupStart.Next.Next
	groupStart.Next.Next = groupStart
	groupStart.Next = groupEnd
	// fmt.Println(dummyHead.Val, curr.Val, groupStart.Val, groupEnd.Val, curr.Next.Next.Val)
	// fmt.Println(dummyHead.Next.Val, dummyHead.Next.Next.Val, dummyHead.Next.Next.Next.Val, dummyHead.Next.Next.Next.Next.Val)
	for groupEnd != nil && groupEnd.Next != nil {
		curr = groupStart
		groupStart = groupEnd
		curr.Next = groupStart.Next
		groupEnd = groupStart.Next.Next
		groupStart.Next.Next = groupStart
		groupStart.Next = groupEnd
	}
	// if groupEnd != nil {
	//     fmt.Println(groupEnd.Val)
	// }
	return dummyHead.Next
}