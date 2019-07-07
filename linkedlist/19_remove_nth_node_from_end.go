package linkedlist

/*
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	q := head
	for i := 0; i < n; i++ {
		q = q.Next
	}
	if q == nil {
		return head.Next
	}
	for q.Next != nil {
		q = q.Next
		p = p.Next
	}
	p.Next = p.Next.Next
	return head
}

// 第一想法是遍历一遍，得到第一个位置需要delete，然后再从头遍历一遍删除
// 那么如何遍历一次呢，那么遍历到最后才知道第一位要delete，那么可以记录每一个位置的指针，建立以array
// 其实并不是每一个位置都需要记录，我们只需要记录到最后后相差的那个位置的指针，所以双指针可解
