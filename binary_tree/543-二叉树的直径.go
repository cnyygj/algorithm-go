package binary_tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	// 空树，或者单节点的树，都返回0，单节点没有最长路径
	if root == nil || root.Left == nil && root.Right == nil{
		return 0
	}

	maxLen := -1 * math.MaxInt64
	var backTravel func(root *TreeNode) int
	backTravel = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftMaxLen := backTravel(root.Left)
		rightMaxLen := backTravel(root.Right)

		if leftMaxLen + rightMaxLen > maxLen {
			maxLen = leftMaxLen + rightMaxLen
		}

		return max(leftMaxLen, rightMaxLen) + 1
	}

	backTravel(root)
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
