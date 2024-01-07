package binary_tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

思路：
先中序遍历，得到一个升序的有序数组，两个元素之差的绝对值的最小值，答案一定是相邻两个元素之差的最小值
所以在遍历节点时，只需要判断当前节点和其前面节点的差值是否是最小值即可
 */
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minAbs := math.MaxInt64
	var lastNum int
	isFirst := true

	update := func(num int) {
		if !isFirst && abs(num - lastNum) < minAbs {
			minAbs = abs(num - lastNum)
		}
		lastNum = num
		isFirst = false
	}

	var midTravel func(root *TreeNode)
	midTravel = func(root *TreeNode) {
		if root == nil {
			return
		}

		midTravel(root.Left)
		update(root.Val)
		midTravel(root.Right)
	}

	// 中序遍历树, 得到有序数组
	midTravel(root)

	return minAbs
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
