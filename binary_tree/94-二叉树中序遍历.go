package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var tree []int
func inorderTraversal(root *TreeNode) []int {
	tree = []int{}
	midTravel(root)
	return tree
}

func midTravel(root *TreeNode) {
	if root == nil {
		return
	}

	midTravel(root.Left)
	tree = append(tree, root.Val)
	midTravel(root.Right)
}

/**
可以不需要定义全局变量
func preorderTraversal(root *TreeNode) []int {
    var arr []int
    var travel func(root *TreeNode)
    travel = func(root *TreeNode) {
        if root == nil {
            return
        }
        travel(root.Left)
        arr = append(arr, root.Val)
        travel(root.Right)
    }
    travel(root)

    return arr
}
 */
