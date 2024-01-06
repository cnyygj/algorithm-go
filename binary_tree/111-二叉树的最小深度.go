package binary_tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

实际上就是求跟节点到叶子节点的最小路径，或者说是树的最小高度，
所谓的叶子节点，就是没有子节点的节点，

根节点的到叶子节点的最小高度，就是根节点左右子树的最小高度+1，左右子树的最小高度又可以分解成它们的左右子树的最小高度+1，
以此类推...
所以可利用递归的方法，求解根节点到叶子节点的最小高度，
 */

/**
方法一：
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	mDepth := math.MaxInt32
	if root.Left != nil {
		mDepth = min(minDepth(root.Left), mDepth)
	}
	if root.Right != nil {
		mDepth = min(minDepth(root.Right), mDepth)
	}

	return 1 + mDepth
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**

方法二：与方法一相似，都是算左右子树最小高度
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    // 左子树为空，算右子树最小高度
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    // 右子树为空，算左子树最小高度
    if root.Right == nil {
        return minDepth(root.Left) + 1
    }

    // 左右子树不为空，取其最小值
    return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
 */

/**
方法三：广度优先遍历
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    var queue []*TreeNode
    queue = append(queue, root)

    depth := 0
    for len(queue) > 0 {
        size := len(queue)
        depth++
        for i := 0; i < size; i++ {
            temp := queue[0]
            queue = queue[1:]

            if temp.Left == nil && temp.Right == nil {
                return depth
            }

            if temp.Left != nil {
                queue = append(queue, temp.Left)
            }
            if temp.Right != nil {
                queue = append(queue, temp.Right)
            }
        }
    }

    return depth
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

 */
