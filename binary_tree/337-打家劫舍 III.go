package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }


小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。

除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。

思路：
我们可以用 f(o) 表示选择 o 节点的情况下，o 节点的子树上被选择的节点的最大权值和；g(o) 表示不选择 o 节点的情况下，o 节点的子树上被选择的节点的最大权值和；l 和 r 代表 o 的左右孩子。

当 o 被选中时，o 的左右孩子都不能被选中，故 o 被选中情况下子树上被选中点的最大权值和为 l 和 r 不被选中的最大权值和相加，即 f(o)=g(l)+g(r)。
当 o 不被选中时，o 的左右孩子可以被选中，也可以不被选中。对于 o 的某个具体的孩子 x，它对 o 的贡献是 x 被选中和不被选中情况下权值和的较大值。故 g(o)=max{f(l),g(l)} + max{f(r),g(r)}。
 */
func rob(root *TreeNode) int {
	sum := dfs(root)
	return max(sum[0], sum[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	left := dfs(node.Left)
	right := dfs(node.Right)

	selected := node.Val + left[1] + right[1]
	noselected := max(left[0], left[1]) + max(right[0], right[1])
	return []int{selected, noselected}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
