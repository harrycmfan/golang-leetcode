package tree

import "fmt"

/*
Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := []string{}
	dfs(root, fmt.Sprintf("%d", root.Val), &res)
	return res
}

func dfs(root *TreeNode, path string, res *[]string) {
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
	}
	if root.Left != nil {
		dfs(root.Left,fmt.Sprintf("%s->%d", path, root.Left.Val), res)
	}
	if root.Right != nil {
		dfs(root.Right,fmt.Sprintf("%s->%d", path, root.Right.Val), res)
	}
}

// func binaryTreePaths(root *TreeNode) []string {
//     if root == nil {
//         return []string{}
//     }
//     var res []string
//     var hasChild bool
//     if root.Left != nil {
//         leftRes := binaryTreePaths(root.Left)
//         for _, l := range leftRes {
//             res = append(res, fmt.Sprintf("%d->%s", root.Val, l))
//         }
//         hasChild = true
//     }
//     if root.Right != nil {
//         rightRes := binaryTreePaths(root.Right)
//         for _, r := range rightRes {
//             res = append(res, fmt.Sprintf("%d->%s", root.Val, r))
//         }
//         hasChild = true
//     }
//     if !hasChild {
//         res = append(res, fmt.Sprintf("%d", root.Val))
//     }
//     return res
// }



