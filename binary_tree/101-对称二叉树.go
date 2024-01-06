package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

所谓的对称，就是两棵树互为镜像
所谓的镜像就是：根节点的值相同，同时，一棵树的左子树和另一棵树的右子树互为镜像
可以用递归的方式实现
 */
func isSymmetric(root *TreeNode) bool {
	return checkNode(root, root)
}

func checkNode(l, r *TreeNode) bool {
	if l == nil  && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return checkNode(l.Left, r.Right) && checkNode(l.Right, r.Left)
}
