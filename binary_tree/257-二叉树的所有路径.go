package binary_tree

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var traces []string
	if root == nil {
		return traces
	}

	var track []string
	var travel func (root *TreeNode)
	travel = func (root *TreeNode) {
		if root == nil {
			return
		}
		// 叶子结点
		if root.Left == nil && root.Right == nil {
			// 路径加上叶子节点（选择）
			track = append(track, strconv.Itoa(root.Val))
			trackStr := strings.Join(track, "->")
			traces = append(traces, trackStr)
			// 归纳完后，把当前叶子节点去掉（撤销选择）
			track = track[0: len(track)-1]
			return
		}

		// 遍历到非叶子节点，路径加上（选择）
		track = append(track, strconv.Itoa(root.Val))
		if root.Left != nil {
			travel(root.Left)
		}
		fmt.Println("root.Val", root.Val, track)
		if root.Right != nil {
			travel(root.Right)
		}

		// 非叶子节点遍历完左右子树后，把这个节点从路径上去掉（撤销选择）
		track = track[0: len(track)-1]

	}
	travel(root)
	return traces
}
