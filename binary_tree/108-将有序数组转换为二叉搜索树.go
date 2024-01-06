package binary_tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

题目要求：将一个升序的数组，转换为一棵 高度平衡 二叉搜索树

高度平衡：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1
二叉搜索树：左子树所有节点的值都小于根节点的值，右子树所有节点的值都大于根节点的值

我们知道，二叉搜索树的中序遍历是升序序列，但是给定二叉搜索树的中序遍历，又不可能唯一确认二叉搜索树

怎么做？

解题思路：
1. 取中间元素作为根节点
2. 递归构建左子树和右子树

因为数组是有序的，所以中间元素左边的所有的元素都比它小，其右边的所有元素都比它大，满足二叉搜索树的条件

同时，选择中间数字作为二叉搜索树的根节点，这样分给左右子树的数字个数相同或只相差 111，可以使得树保持平衡：
	- 如果数组长度是奇数，则根节点的选择是唯一的，
	- 如果数组长度是偶数，则可以选择中间位置左边的数字作为根节点或者选择中间位置右边的数字作为根节点，

选择不同的数字作为根节点则创建的平衡二叉搜索树也是不同的。这个是允许的。

 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return buildTree(0, len(nums)-1, nums)
}

func buildTree(left, right int, nums []int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right - left) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildTree(left, mid-1, nums)
	root.Right = buildTree(mid+1, right, nums)
	return root
}
