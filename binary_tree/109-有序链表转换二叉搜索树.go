package binary_tree

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。

思路：
由于单链表中的元素已经是升序，我们可以将链表中的元素转换为切片，改切片可以认为是二叉搜索树的中序遍历
将有序数组构造成平衡二叉树
每次取中间的数作为根节点, 由中间的数将数组分为左右两部分分别构建左右子树，这样构造的数的高度差不会超过1
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return buildTee(0, len(nums)-1, nums)
}

func buildTee(left, right int, nums []int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right - left) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildTee(left, mid-1, nums)
	root.Right = buildTee(mid+1, right, nums)

	return root
}
