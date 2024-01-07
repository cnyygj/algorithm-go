package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

思路：
先中序遍历，得到有序数组，然后遍历数组，得到出现次数最多的数
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var midArr []int
	var midTravel func(root *TreeNode)
	midTravel = func(root *TreeNode) {
		if root == nil {
			return
		}

		midTravel(root.Left)
		midArr = append(midArr, root.Val)
		midTravel(root.Right)
	}

	// 中序遍历树, 得到有序数组
	midTravel(root)

	return findMostFrequent(midArr)
}

func findMostFrequent(arr []int) []int {
	var mostFrequent []int
	if len(arr) == 0 {
		return mostFrequent
	}

	numCount := make(map[int]int)
	var maxCount int
	for _, num := range arr {
		numCount[num]++
		if numCount[num] > maxCount {
			maxCount = numCount[num]
		}
	}

	for num, count := range numCount {
		if count == maxCount {
			mostFrequent = append(mostFrequent, num)
		}
	}

	return mostFrequent
}

/**
第二种方法：

使用哈希表，遍历二叉树，统计每个数字出现的次数，最后找到最大次数的数字。
时间复杂度：O(n)，其中 n 是二叉搜索树的节点数。
空间复杂度：O(n)。即递归的栈空间的空间代价。

func findMode(root *TreeNode) (answer []int) {
    var base, count, maxCount int

    update := func(x int) {
        if x == base {
            count++
        } else {
            base, count = x, 1
        }
        if count == maxCount {
            answer = append(answer, base)
        } else if count > maxCount {
            maxCount = count
            answer = []int{base}
        }
    }

    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        update(node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return
}
 */

