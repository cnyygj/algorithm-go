package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

所谓的平衡二叉树：就是二叉树的任意一个节点的左右子树的高度差的绝对值不超过1
使用递归的方式，求解二叉树的高度，然后判断是否平衡
 */

/**
方法一
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(heightTree(root.Left) - heightTree(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
	/**
	这里需要注意的是：不能分开写上面的条件，如：
	if abs(heightTree(root.Left) - heightTree(root.Right)) <= 1 {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right)

	原因是：即使当前左右子树是平衡的，但是可能其子树不是平衡的，所以需要继续判断其左右子树。
	 */
}

// 就算树的高度，也是使用递归的方式
func heightTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(heightTree(root.Left), heightTree(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

/**
方法二：
基于方法一可以知道，每次判断当前节点树是否平衡的时候，都需要递归算出当前节点树的高度，
其实没必要，只需要先对树递归遍历一次算出每个节点的高度，然后存储到root.Val中即可，判断树是否平衡时，只需要判断该节点左节点的Val和右节点的Val是否满足平衡二叉树的条件即可。

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    heightTree(root)
    return checkBalance(root)
}

func checkBalance(root *TreeNode) bool {
    if root.Left == nil && root.Right == nil {
        return true
    }

	// 需要把一些左节点或右节点是否为空枚举判断
    if (root.Right == nil && root.Left.Val > 1) || (root.Left == nil && root.Right.Val > 1) {
        return false
    }
    if (root.Right == nil && root.Left.Val <= 1) || (root.Left == nil && root.Right.Val <= 1) {
        return true
    }

    return abs(root.Left.Val - root.Right.Val) <= 1 && checkBalance(root.Left) && checkBalance(root.Right)
}

func heightTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

	// 把当前节点的高度存储到root.Val中
    root.Val = 1 + max(heightTree(root.Left), heightTree(root.Right))
    return root.Val
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -1 * a
    }
    return a
}
 */
