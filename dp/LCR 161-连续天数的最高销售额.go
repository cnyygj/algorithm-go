package dp


/**
某公司每日销售额记于整数数组 sales，请返回所有 连续 一或多天销售额总和的最大值。

要求实现时间复杂度为 O(n) 的算法。

示例 1：
输入：sales = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：[4,-1,2,1] 此连续四天的销售总额最高，为 6。

示例 2：
输入：sales = [1,-2,1,-2,4,-4]
输出：2

 */

/**
动态规划
时间复杂度 O(n)
空间复杂度 O(n)
dp[i] 表示以 sales[i] 为结尾的连续子数组的最大和
dp[i] = max(dp[i-1] + sales[i], sales[i])
维护一个全局变量 maxAmount，表示所有 dp[i] 的最大值
遍历数组，更新 maxAmount
最终返回 maxAmount
 */
func maxSales(sales []int) int {
	n := len(sales)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return sales[0]
	}

	dp := make([]int, n)
	dp[0] = sales[0]

	maxAmount := dp[0]
	for j := 1; j < n ; j++ {
		dp[j] = max(dp[j-1] + sales[j], sales[j])
		maxAmount = max(maxAmount, dp[j])
	}

	return maxAmount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
