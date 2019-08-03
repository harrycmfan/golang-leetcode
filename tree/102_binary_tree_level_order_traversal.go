package tree

/*

Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
 */


func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	bfs(root, 0, &res)
	return res
}

func bfs(root *TreeNode, level int, res *[][]int) {
	if level >= len(*res) {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], root.Val)
	if root.Left != nil {
		bfs(root.Left, level+1, res)
	}
	if root.Right != nil {
		bfs(root.Right, level+1, res)
	}
}