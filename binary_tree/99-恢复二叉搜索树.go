package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。

输入：root = [1,3,null,null,2]
输出：[3,1,null,null,2]
解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。

输入：root = [3,1,4,null,null,2]
输出：[2,1,4,null,null,3]
解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
 */

/**
方法一：
1. 中序遍历二叉搜索树，将所有节点的值存储在一个数组中，如果是二叉搜索树，那么数组应该是升序的，显然，交换完两个节点后，数组的i, j 位置，一定会存在 ai > ai+1 ， aj < aj+1
2. 遍历数组，找到两个被错误地交换的节点
3. 交换这两个节点的值
 */
func recoverTree(root *TreeNode)  {

	var nums []int
	var midTravel func(root *TreeNode)
	midTravel = func(root *TreeNode) {
		if root == nil {
			return
		}
		midTravel(root.Left)
		nums = append(nums, root.Val)
		midTravel(root.Right)
	}

	midTravel(root)
	x, y := findSwaped(nums)
	ajustTree(root, 2, x, y)

}

func findSwaped(nums []int) (int, int) {
	a, b := -1, -1
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i+1] < nums[i] {
			b = i+1
			if a == -1 {
				a = i
			} else {
				// 因为只有两个节点，所以当两个节点都找到时，直接退出循环
				break
			}
		}
	}

	return nums[a], nums[b]
}

func ajustTree(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}

	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}

		// 增加count参数，减少递归次数
		count--
		if count == 0 {
			return
		}
	}

	ajustTree(root.Left, count, x, y)
	ajustTree(root.Right, count, x, y)
}

/**
方法二：
1. 因为二叉搜索树的中序遍历是升序的，所以中序遍历二叉搜索树，得到的是一个升序
2. 遍历二叉搜索树，比较当前节点和前一个节点，如果当前节点小于前一个节点，说明当前节点和前一个节点被错误地交换了
3. 交换这两个节点
 */

func recoverTree(root *TreeNode) {
	pre = nil
	node1 = nil
	node2 = nil
	recoverTreeHelp(root)
	node1.Val, node2.Val = node2.Val, node1.Val
}

var pre *TreeNode
var node1 *TreeNode
var node2 *TreeNode

func recoverTreeHelp(root *TreeNode)  {

	if root == nil{
		return
	}

	recoverTreeHelp(root.Left)

	if pre == nil {
		pre = &TreeNode{Val:root.Val}
	} else if pre.Val > root.Val {
		if node1 == nil {
			node1 = pre
		}
		node2 = root
	}

	pre = root

	recoverTreeHelp(root.Right)

}
