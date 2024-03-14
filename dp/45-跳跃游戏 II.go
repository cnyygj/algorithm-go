package dp

/**
给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2
 */

/**
思路：动态规划
dp[i] 表示到达第 i 个位置的最小跳跃次数
dp[i] = min(dp[i], dp[j]+1)，条件：j ∈ [0, i-1]，j + nums[j] >= i
初始化：dp[0] = 0
结果：dp[n-1]
 */
func jump(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = n
	}

	for i := 1; i < n; i++ {
		for j := i-1; j >= 0; j-- {
			if (j + nums[j]) >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
思路：贪心
每次选择最大的跳跃步数，同时更新最大跳跃步数
当走到上次最大的跳跃步数的边界时，跳跃步数加一，同时更新边界值
 */
func jump2(nums []int) int {
	numLen := len(nums)
	currentRight := 0
	maxRight := 0
	var count = 0
	for i := 0; i < numLen-1; i++ {

		right := i + nums[i]
		if right > maxRight {
			maxRight = right
		}

		if i == currentRight {
			currentRight = maxRight
			count++
		}
	}
	return count
}
