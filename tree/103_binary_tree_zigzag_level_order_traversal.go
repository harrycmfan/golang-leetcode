package tree

/*
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]

 */


func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	bfs(root, 1, &res)
	return res
}

func bfs(root *TreeNode, height int, res *[][]int) {
	if len(*res) < height { // to prevent out of index
		*res = append(*res, []int{})
	}
	if height % 2 == 1 {
		(*res)[height-1] = append((*res)[height-1], root.Val)
	} else {
		(*res)[height-1] = append([]int{root.Val}, (*res)[height-1]...)
	}

	if root.Left != nil {
		bfs(root.Left, height + 1, res)
	}
	if root.Right != nil {
		bfs(root.Right, height + 1, res)
	}
}