package linkedlist

/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	copyLists := lists
	for len(copyLists) > 1 {
		result := []*ListNode{}
		for i:= 0; i < len(copyLists); i+=2 {
			if i+1 < len(copyLists) {
				result = append(result, mergeTwoLists(copyLists[i], copyLists[i+1]))
			} else {
				result = append(result, copyLists[i])
			}
		}
		copyLists = result
	}
	return copyLists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	dummyHead := &ListNode{}
	curr := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	for l1 != nil {
		curr.Next = l1
		l1 = l1.Next
		curr = curr.Next
	}
	for l2 != nil {
		curr.Next = l2
		l2 = l2.Next
		curr = curr.Next
	}
	return dummyHead.Next
}