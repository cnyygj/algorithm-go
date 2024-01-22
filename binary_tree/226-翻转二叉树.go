package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

如果当前遍历到的节点 root 的左右两棵子树都已经翻转，
那么我们只需要交换两棵子树的位置，即可完成以 root 为根节点的整棵子树的翻转。

因此可以递归从叶子节点先开始翻转

 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

/**
方法二：
也可以自上而下翻转，相当于两个树，一个从左遍历，一个从右遍历，然后交换

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    root.Left, root.Right = revTree(root.Left, root.Right)

    return root
}

func revTree(leftNode, rightNode *TreeNode) (*TreeNode, *TreeNode) {
    if leftNode == nil && rightNode == nil {
        return nil, nil
    }

    if leftNode == nil {
        leftNode = &TreeNode{Val: rightNode.Val}
 		// 需要对rightNode的左右节点进行翻转，并发翻转后的左右节点重新赋值给leftNode的左右节点
        leftNode.Left, leftNode.Right = revTree(rightNode.Left, rightNode.Right)
        rightNode = nil
        return leftNode, rightNode
    }

    if rightNode == nil {
        rightNode = &TreeNode{Val: leftNode.Val}
		// 需要对leftNode的左右节点进行翻转，并发翻转后的左右节点重新赋值给rightNode的左右节点
        rightNode.Left, rightNode.Right = revTree(leftNode.Left, leftNode.Right)
        leftNode = nil
        return leftNode, rightNode
    }

    temp := leftNode.Val
    leftNode.Val = rightNode.Val
    rightNode.Val = temp

    leftNode.Left, rightNode.Right = revTree(leftNode.Left, rightNode.Right)
    leftNode.Right, rightNode.Left = revTree(leftNode.Right, rightNode.Left)
    return leftNode, rightNode
}
 */
