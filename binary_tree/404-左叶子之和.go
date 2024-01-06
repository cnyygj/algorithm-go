package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

给定二叉树的根节点 root ，返回所有左叶子之和。
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftCount int
	var sumLeft, sumRight func(root *TreeNode)
	// 针对左子树的处理
	sumLeft = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 只在左子树中存在的左叶子节点才做计算
		if root.Left == nil && root.Right == nil {
			leftCount += root.Val
			return
		}

		// 左子树中可能存在左节点和右节点
		sumLeft(root.Left)
		sumRight(root.Right)
	}
	// 针对右子树的处理
	sumRight = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 右子树右叶子节点可以跳过不统计
		if root.Left == nil && root.Right == nil {
			return
		}

		// 同样，右子树中也可能存在左节点和右节点
		sumLeft(root.Left)
		sumRight(root.Right)
	}

	// 计算左子树
	sumLeft(root.Left)
	// 计算右子树
	sumRight(root.Right)

	return leftCount
}
