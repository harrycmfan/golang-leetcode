package linkedlist

/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var previousNode *ListNode
	curr := head
	for curr != nil {
		NextToTravel := curr.Next
		curr.Next = previousNode
		previousNode = curr
		curr = NextToTravel
	}
	return previousNode
}