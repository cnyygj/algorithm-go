package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

输入：n = 3
输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

如果我们枚举根节点的值为 iii，那么根据二叉搜索树的性质我们可以知道左子树的节点值的集合为 [1…i−1][1 \ldots i-1][1…i−1]，右子树的节点值的集合为 [i+1…n][i+1 \ldots n][i+1…n]。而左子树和右子树的生成相较于原问题是一个序列长度缩小的子问题，因此我们可以想到用回溯的方法来解决这道题目。
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return getAllTree(1, n)
}

// 回溯法找全排列
func getAllTree(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var trees []*TreeNode

	for i := start; i <= end; i++ {
		leftTree := getAllTree(start, i-1)
		rightTree := getAllTree(i+1, end)

		for _, left := range leftTree {
			for _, right := range rightTree {
				node := &TreeNode{Val: i}
				node.Left = left
				node.Right = right
				trees = append(trees, node)
			}
		}
	}

	return trees
}
