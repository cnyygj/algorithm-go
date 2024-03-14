package dp

/**
给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。

示例 1：

输入：cost = [10,15,20]
输出：15
解释：你将从下标为 1 的台阶开始。
- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
总花费为 15 。
 */

/*
 动态规划
 时间复杂度：O(n)
 空间复杂度：O(n)

 dp[i] 表示到达第 i 个台阶时的最低花费
 dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
 初始化：dp[0] = cost[0], dp[1] = cost[1]
 */
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[n-1]
	}

	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
