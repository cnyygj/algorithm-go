package dp

/**
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

输入: rowIndex = 3
输出: [1,3,3,1]

*/


func getRow(rowIndex int) []int {
	var preRow []int
	for i := 0; i <= rowIndex; i++ {
		if i == 0 {
			preRow = []int{1}
		} else {
			tmp := make([]int, i+1)
			tmp[0] = 1
			tmp[i] = 1
			for j := 1; j < i; j++ {
				tmp[j] = preRow[j-1] + preRow[j]
			}
			preRow = tmp
		}

		if i == rowIndex {
			break
		}
	}

	return preRow
}
