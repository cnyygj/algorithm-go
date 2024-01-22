package dp

/**
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

输入: numRows = 1
输出: [[1]]

*/

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)

	for i := range dp {
		dp[i] = make([]int, i+1)
		dp[i][0] = 1
		dp[i][i] = 1

		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}

	return dp
}
