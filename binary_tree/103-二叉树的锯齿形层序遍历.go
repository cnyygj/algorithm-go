package binary_tree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]

*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	var nodes [][]int
	if root == nil {
		return nodes
	}

	var queue []*TreeNode
	queue = append(queue, root)


	var collectNode func(queue []*TreeNode, depth int) []*TreeNode
	collectNode = func(queue []*TreeNode, depth int) []*TreeNode {
		size := len(queue)
		var temp []int
		var node *TreeNode
		for i := 0; i < size; i++ {
			node = queue[0]
			queue = queue[1:]


			fmt.Printf("node: %v, depth: %v, depth/2: %v\n", node.Val, depth, depth % 2)

			temp = append(temp, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		// 奇数的层数，把子节点的顺序翻转即可
		if depth % 2 == 0 {
			for i, n := 0, len(temp)-1; i <= n; i, n = i+1, n-1{
				temp[i], temp[n] = temp[n], temp[i]
			}
		}
		nodes = append(nodes, temp)
		return queue
	}

	depth := 1
	for len(queue) > 0 {
		queue = collectNode(queue, depth)
		depth++
	}

	return nodes
}
