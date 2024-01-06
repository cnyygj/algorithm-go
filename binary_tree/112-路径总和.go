package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum

题目要求：必须要是根节点到叶子节点，也就是说不能从中间节点开始。或者是说，路径的终点是非叶子节点

思路：
	1. 递归遍历整棵树
	2. 递归函数的参数：当前节点，剩余的目标值
	3. 如果当前节点是叶子节点，判断剩余目标值是否为0，如果为0，返回true，否则返回false
	4. 如果当前节点不是叶子节点，递归遍历当前节点的左子树和右子树
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if targetSum - root.Val == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	if hasPathSum(root.Left, targetSum - root.Val) {
		return true
	}

	if hasPathSum(root.Right, targetSum - root.Val) {
		return true
	}

	return false
}

/**
方法二：
实际上也是求 两个点之间的路径和，只不过这个路径和是经过根节点和叶子节点的路径和
使用广度优先遍历方法遍历整棵树，找到路径和等于目标值的路径，并返回true，否则返回false

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    nodes := []*TreeNode{root}
    values :=[]int{root.Val}

    for len(nodes) > 0 {
        node := nodes[0]
        nodes = nodes[1:]
        val := values[0]
        values = values[1:]
        if node.Left == nil && node.Right == nil && targetSum == val {
            return true
        }

        if node.Left != nil {
            nodes = append(nodes, node.Left)
            values = append(values, node.Left.Val+val)
        }

        if node.Right != nil{
            nodes = append(nodes, node.Right)
            values = append(values, node.Right.Val+val)
        }
    }


    return false

}
 */
