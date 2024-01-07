package binary_tree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
思路：就是遍历root的子树，判断子树是否与subRoot相同，相同就返回true，否则继续遍历左右子树
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	var findSubTree func(root, subRoot *TreeNode) bool
	findSubTree = func(root, subRoot *TreeNode) bool {
		if root == nil {
			return false
		}

		if root.Val == subRoot.Val && checkSameTree(root, subRoot) {
			return true
		}

		if findSubTree(root.Left, subRoot) {
			return true
		}

		return findSubTree(root.Right, subRoot)
	}

	return findSubTree(root, subRoot)
}

func checkSameTree(subTree, tragetTree *TreeNode) bool {
	if subTree == nil && tragetTree == nil {
		return true
	}

	if subTree == nil || tragetTree == nil {
		return false
	}

	return subTree.Val == tragetTree.Val && checkSameTree(subTree.Left, tragetTree.Left) && checkSameTree(subTree.Right, tragetTree.Right)
}

/**
代码优化：

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil || subRoot == nil {
        return false
    }

    return checkSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func checkSameTree(subTree, tragetTree *TreeNode) bool {
    if subTree == nil && tragetTree == nil {
        return true
    }

    if subTree == nil || tragetTree == nil {
        return false
    }

    return subTree.Val == tragetTree.Val && checkSameTree(subTree.Left, tragetTree.Left) && checkSameTree(subTree.Right, tragetTree.Right)
}
 */
