package dp

/**
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。



示例 1:
输入: s = "abc", t = "ahbgdc"
输出: true

示例 2:
输入: s = "axc", t = "ahbgdc"
输出: false
 */


/**
方法一：双指针

时间复杂度：O(n)
空间复杂度：O(1)

思路：
1. 遍历t，如果t[right] == s[left]，则left++
2. 最后判断left是否等于s的长度
3. 返回结果
 */
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	var left, right int
	var preLeftStr byte
	var preRightStr byte

	for right < len(t) {
		leftStr := s[left]
		rightStr := t[right]

		// 针对 t字符串中出现相同字符串的情况做处理，比如：s = "leet", t = "leeeeeeeeeeeetcode"
		if preLeftStr > 0 && preRightStr > 0 && rightStr == preRightStr && leftStr != preLeftStr {
			right++
		} else {
			if leftStr == rightStr {
				left++
			}

			if left == len(s) {
				return true
			}

			preLeftStr = leftStr
			preRightStr = rightStr
			right++
		}
	}

	return false
}

/**
方法二：动态规划

时间复杂度：O(n*m)
空间复杂度：O(n*m)

思路：
1. 定义dp[i][j]表示s的前i个字符和t的前j个字符的公共子序列的长度
2. 初始化dp[0][0] = 0，表示空字符串和空字符串的公共子序列的长度为0
3. 遍历s和t，如果s[i-1] == t[j-1]，说明这个字符一定是在公共子序列中，那么公共子序列的长度应该再加1，则dp[i][j] = dp[i-1][j-1] + 1
4. 否则，说明s[i-1] 和 t[j-1] 这两个字符至少有一个不在公共子序列中，dp[i][j] = max(dp[i-1][j], dp[i][j-1])
5. 最后判断dp[n][m]是否等于s的长度
6. 返回结果
 */
func isSubsequence2(s string, t string) bool {
	n1 := len(s)
	n2 := len(t)
	dp := make([][]int, n1+1)
	for i, _ := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2] == n1
}
