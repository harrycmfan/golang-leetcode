package tree

/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var stack []*TreeNode
	a := []int{}
	//stack = append(stack, root)
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop出来,就是如果在stack中取出来的必须要pop，不然就会遍历两次
		a = append(a, root.Val)
		root = root.Right // 主动走到右边
		//fmt.Println(a)
	}
	return a
}

// func inorderTraversal(root *TreeNode) []int {
//     if root == nil {
//         return []int{}
//     }
//     a := []int{}
//     a = append(a, inorderTraversal(root.Left)...)
//     a = append(a, root.Val)
//     a = append(a, inorderTraversal(root.Right)...)
//     return a
// }