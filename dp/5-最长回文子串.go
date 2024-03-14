package dp

/**
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的最长回文子串。
示例 2：

输入：s = "cbbd"
输出："bb"
 */

/**
动态规划
dp[i][j] 表示 s[i...j] 是否是回文串
dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
状态转移方程：
if s[i] == s[j] && (j - i == 1 || dp[i+1][j-1]) {
    dp[i][j] = true
}
时间复杂度：O(n^2)
空间复杂度：O(n^2)
 */
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	result := s[0:1]

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for i := n-2; i >= 0 ; i-- {
		for j := i+1; j < n; j++ {
			if s[i] != s[j] {
				continue
			}

			// 判断一下相邻的两个数是否是回文
			if j - i == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1]
			}

			// 若果当前 s[i: j+1]是回文，就计算最长回文
			if dp[i][j] && len(result) < j-i+1 {
				result = s[i:j+1]
			}
		}
	}

	return result

}
