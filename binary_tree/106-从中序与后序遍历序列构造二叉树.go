package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。


输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]

输入：inorder = [-1], postorder = [-1]
输出：[-1]

*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0{
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	i := 0
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}

	// 注意：不管是左子树还是右子树，其对应的中序遍历和后序遍历的数组的长度都是一样的
	// 因此，在确定root节点在中序遍历数组的位置后，该位置左边为左子树，右边为右子树，计算出左右子树的长度后，要对后序遍历数组做切割
	root.Left = buildTree(inorder[:i], postorder[:len(inorder[:i])])
	root.Right = buildTree(inorder[i+1:], postorder[len(inorder[:i]):len(postorder)-1])



	return root
}
