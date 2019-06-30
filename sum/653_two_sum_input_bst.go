package sum

/*Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:

Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

Output: True


Example 2:

Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

Output: False
 */



//Definition for a binary tree node.
type TreeNode struct {
 	Val int
    Left *TreeNode
    Right *TreeNode
}

var myMap = map[int]bool{}

func findTarget(root *TreeNode, k int) bool {
	defer func(){
		myMap = map[int]bool{}
	}()
	return checkMap(root, k)
}

func checkMap(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	_, ok := myMap[root.Val]
	if ok {
		return true
	}
	myMap[k - root.Val] = true
	return checkMap(root.Left, k) || checkMap(root.Right, k)
}


// 二叉树中序遍历得到一个有序数组，然后有序数组通过前后指针遍历结果
// 任意方式遍历，通过map来保存结果
