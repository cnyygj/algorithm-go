package dp

/**
给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。



示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 */

func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	maxRight := 0

	for i := 0; i < n-1; i++ {

		// 如果当前能跳跃的步数为0，且当前位置也是截止目前为止最大的跳跃位置（如：[1,0,1,0]）,直接返回false
		if nums[i] == 0 && i == maxRight {
			return false
		}

		// 更新最大跳跃边界
		right := i + nums[i]
		if right > maxRight {
			// 如果最大的跳跃边界刚好在最后一个下标的位置，或者超过最后一个下标的位置，返回true
			if right >= n-1 {
				return true
			}
			maxRight = right
		}
	}

	return false
}

/**
方法二：动态规划
dp[i]：从[0,i]的任意一点处出发，你最大可以跳跃到的位置。

例如 nums=[2,3,1,1,4]中:
dp[0]=2, dp[1]=4, dp[2]=4, dp[3]=4, dp[4]=8（其实我们没有必要去讨论最后一个下标，因为从最后一个下标出发一定可以到最后一个）

因此递推公式
dp[i] = max(dp[i-1],nums[i]+i)

由dp[i]的定义可以知道，dp[0]=nums[0]

当dp[i] == i时，说明从[0,i]的任意一点出发，最大可以跳跃到的位置就是i，不能再仅需前进了，直接返回false。
当dp[i] >= n-1时，说明从[0,i]的任意一点出发，最大可以跳跃到的位置就是最后一个下标，直接返回true。

*/
func canJump2(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+i)
		if dp[i] >= n-1 {
			return true
		}

		// 当dp[i] == i时，说明从[0,i]的任意一点出发，最大可以跳跃到的位置就是i，不能再仅需前进了，直接返回false。
		if dp[i] == i {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

