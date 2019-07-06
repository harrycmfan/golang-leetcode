package linkedlist

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
 */

 type ListNode struct {
     Val int
 	Next *ListNode
 }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	sum := l1.Val + l2.Val
	advance := 0
	if sum >= 10 {
		sum = sum - 10
		advance = 1
	}
	result.Val = sum
	currentList := result
	for l1.Next != nil && l2.Next != nil {
		l1 = l1.Next
		l2 = l2.Next
		currentList.Next = &ListNode{}
		currentList = currentList.Next
		sum = l1.Val + l2.Val + advance
		advance = 0
		if sum >= 10 {
			sum = sum - 10
			advance = 1
		}
		currentList.Val = sum
	}
	var remainList *ListNode
	if l1.Next != nil {
		remainList = l1.Next
	} else {
		remainList = l2.Next
	}
	for remainList != nil {
		currentList.Next = &ListNode{}
		currentList = currentList.Next
		sum = remainList.Val + advance
		advance = 0
		if sum >= 10 {
			sum = sum - 10
			advance = 1
		}
		currentList.Val = sum
		remainList = remainList.Next
	}
	if advance == 1 {
		currentList.Next = &ListNode{Val:1}
	}

	return result
}