package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

输入: preorder = [-1], inorder = [-1]
输出: [-1]

思路：
1. 先序遍历的第一个节点是根节点
2. 在中序遍历中找到根节点的位置，根节点左边的为左子树，右边的为右子树
3. 递归构建左右子树
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	i := 0
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	// len(inorder[:i])+1 由中序遍历确定左子树在前序遍历数组中的长度
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
