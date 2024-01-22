package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }

给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

示例 1:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。

示例 2：
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 只要p, q 任何一个节点是当前节点，直接返回当前节点，
	// 因为p, q 都是root的子节点，所以p, q的最近公共祖先节点，要么是当前节点，要么是
	// 当前节点的子节点。
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)

	if leftNode != nil && rightNode != nil {
		return root
	}

	if leftNode == nil && rightNode == nil {
		return nil
	}

	if leftNode == nil {
		return rightNode
	} else {
		return leftNode
	}
}

/**
方法二：
题目说明二叉树是二叉搜索树，那么对于二叉搜索树，如果当前节点值大于p和q的值，说明p和q在当前节点的左子树
如果当前节点值小于p和q的值，说明p和q在当前节点的右子树
如果当前节点值在p和q之间，则当前节点就是最近公共祖先
 */

func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}
