package binary_tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]

思路：
使用回溯法把所有结果都收集起来，然后判断是否符合条件
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var allPath [][]int
	if root == nil {
		return allPath
	}

	var track []int
	var getAllPath func(root *TreeNode, targetSum int)
	getAllPath = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && root.Val == target {
			track = append(track, root.Val)
			tmp := make([]int,len(track),cap(track))
			copy(tmp, track)
			allPath = append(allPath, tmp)
			track = track[: len(track)-1]
			return
		}

		track = append(track, root.Val)
		getAllPath(root.Left, target-root.Val)
		getAllPath(root.Right, target-root.Val)
		track = track[: len(track)-1]
	}

	getAllPath(root, targetSum)
	return allPath
}
