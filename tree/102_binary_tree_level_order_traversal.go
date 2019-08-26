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

type A struct {
	Node *TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	queue := []A{A{
		Node: root,
		Level: 1,
	}}
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		if len(res) < elem.Level {
			res = append(res, []int{})
		}
		res[elem.Level-1] = append(res[elem.Level-1], elem.Node.Val)
		if elem.Node.Left != nil {
			queue = append(queue, A{Node:elem.Node.Left, Level:elem.Level+1})
		}
		if elem.Node.Right != nil {
			queue = append(queue, A{Node:elem.Node.Right, Level:elem.Level+1})
		}
	}
	return res
}