package tree

/*
Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	a := []int{}
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			if root.Left != nil {
				root = root.Left
			} else if root.Right != nil {
				root = root.Right
			} else {
				root = nil
			} // 不断向前
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			a = append(a, root.Val)
			if len(stack) > 0 && stack[len(stack)-1].Left == root { // 说明取出的那个是左值
				root = stack[len(stack)-1].Right
			} else {
				root = nil
			}
		}
		//fmt.Println(stack)
	}

	// a = append(a, postorderTraversal(root.Left)...)
	// a = append(a, postorderTraversal(root.Right)...)
	// a = append(a, root.Val)
	return a
}

// 有一种解法，post是左右中，那么可以变成遍历中右左，类似于preorder，但是先把left push进去，最后倒过来输出

/*
Pre Order Traverse
public List<Integer> preorderTraversal(TreeNode root) {
    List<Integer> result = new ArrayList<>();
    Deque<TreeNode> stack = new ArrayDeque<>();
    TreeNode p = root;
    while(!stack.isEmpty() || p != null) {
        if(p != null) {
            stack.push(p);
            result.add(p.val);  // Add before going to children
            p = p.left;
        } else {
            TreeNode node = stack.pop();
            p = node.right;
        }
    }
    return result;
}
In Order Traverse
public List<Integer> inorderTraversal(TreeNode root) {
    List<Integer> result = new ArrayList<>();
    Deque<TreeNode> stack = new ArrayDeque<>();
    TreeNode p = root;
    while(!stack.isEmpty() || p != null) {
        if(p != null) {
            stack.push(p);
            p = p.left;
        } else {
            TreeNode node = stack.pop();
            result.add(node.val);  // Add after all left children
            p = node.right;
        }
    }
    return result;
}
Post Order Traverse
public List<Integer> postorderTraversal(TreeNode root) {
    LinkedList<Integer> result = new LinkedList<>();
    Deque<TreeNode> stack = new ArrayDeque<>();
    TreeNode p = root;
    while(!stack.isEmpty() || p != null) {
        if(p != null) {
            stack.push(p);
            result.addFirst(p.val);  // Reverse the process of preorder
            p = p.right;             // Reverse the process of preorder
        } else {
            TreeNode node = stack.pop();
            p = node.left;           // Reverse the process of preorder
        }
    }
    return result;
}
 */