package binary_tree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

解题思路：
1、中序遍历，得到有序数组
2、将有序数组构造成平衡二叉树
	2.1、每次取中间的数作为根节点
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
