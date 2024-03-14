package dp

/**
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。

输入：n = 3
 输出：4
 说明: 有四种走法

输入：n = 5
 输出：13

*/

/**
 * 动态规划
 * dp[i] 表示 i 阶台阶有多少种走法
 * dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
 * 初始化：dp[0] = 0, dp[1] = 1, dp[2] = 2, dp[3] = 4
 * 返回 dp[n]
 */

func waysToStep(n int) int {
	if n == 0 {
		return 0
	}

	dp := make(map[int]int)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4

	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
	}

	return dp[n]
}

/**
空间复杂度优化
 */
func waysToStep2(n int) int {
	if n <= 2 {
		return n
	}

	a := 1
	b := 2
	c := 4
	total := c

	for i := 4; i <= n; i++ {
		total = (a + b + c) % 1000000007
		a = b
		b = c
		c = total
	}
	return total
}
