package binary_tree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

给你一棵二叉搜索树，请你返回一棵 平衡后 的二叉搜索树，新生成的树应该与原来的树有着相同的节点值。如果有多种构造方法，请你返回任意一种。
如果一棵二叉搜索树中，每个节点的两棵子树高度差不超过 1 ，我们就称这棵二叉搜索树是 平衡的 。

解题思路：
1、中序遍历，得到有序数组
2、将有序数组构造成平衡二叉树
	2.1、每次取中间的数作为根节点, 由中间的数将数组分为左右两部分分别构建左右子树，这样构造的数的高度差不会超过1
 */
func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var arr []int
	arr = midTravel(root, arr)
	fmt.Println(arr)

	return buildTree(0, len(arr)-1, arr)
}

func midTravel(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}

	arr = midTravel(root.Left, arr)

	arr = append(arr, root.Val)

	arr = midTravel(root.Right, arr)

	return arr
}

func buildTree(left, right int, arr []int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right - left) / 2
	root := &TreeNode{Val: arr[mid]}
	root.Left = buildTree(left, mid - 1, arr)
	root.Right = buildTree(mid + 1, right, arr)

	return root
}
