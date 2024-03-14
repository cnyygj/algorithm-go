package dp

/**
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
 */

/**
思路：
dp[i][j] 表示 word1[0:i] 转换成 word2[0:j] 所使用的最少操作数
dp[i][j] = dp[i-1][j-1] (word1[i] == word2[j])
dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1 (word1[i] != word2[j])
 */
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 0
			} else if i == 0 && j != 0 {
				dp[i][j] = j
			} else if i != 0 && j == 0 {
				dp[i][j] = i
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					// 当word1的第i个字符与word2的第j个字符不相等，即word1[i-1]!=word2[j-1]，存在以下3种情况：
					// 替换：在word1的前i-1个字符和word2的前j-1个字符已经转换完毕的基础上，将word1的第i个字符替换为word2的第j个字符即可，即dp[i][j]=dp[i-1][j-1]+1
					// 插入：在word1的前i个字符和word2的前j-1个字符已经转换完毕的基础上，在word1的后面插入word2的第j个字符即可，即dp[i][j]=dp[i][j-1]+1
					// 删除：在word1的前i-1个字符和word2的前j个字符已经转换完毕的基础上，将word1的第i个字符进行删除即可，即dp[i][j]=dp[i-1][j]+1
					// 由于要求最小操作数，因此dp[i][j]取上述3种情况的较小值
					dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
				}
			}
		}
	}

	return dp[m][n]
}

func min(a, b, c int) int {
	var minNum int
	if a < b {
		minNum = a
	} else {
		minNum = b
	}

	if minNum > c {
		minNum = c
	}

	return minNum
}
