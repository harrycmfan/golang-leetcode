package tree

/*
Given a binary tree, return the preorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]
Follow up: Recursive solution is trivial, could you do it iteratively?
 */

//Definition for a binary tree node.
type TreeNode struct {
    Val int
   	Left *TreeNode
    Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	var a []int
	for len(stack) > 0 {  // 想一想什么屎后进先出 LIFO，你先要printroot本身，然后左，然后右，那么自然是先把右边的压入栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop出来
		a = append(a, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return a
}

// func preorderTraversal(root *TreeNode) []int {
//     if root == nil {
//         return []int{}
//     }
//     a := []int{root.Val}
//     a = append(a, preorderTraversal(root.Left)...)
//     a = append(a, preorderTraversal(root.Right)...)
//     return a
// }