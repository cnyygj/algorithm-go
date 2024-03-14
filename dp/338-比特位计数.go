package dp

/**
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。



示例 1：

输入：n = 2
输出：[0,1,1]
解释：
0 --> 0
1 --> 1
2 --> 10
 */

func countBits(n int) []int {
	var ans []int

	for i := 0; i <= n; i++ {
		if i == 0 {
			ans = append(ans, i)
			continue
		}

		target := i
		var res int
		for target > 0 {
			target &= target-1
			res++
		}
		ans = append(ans, res)
	}


	return ans
}

/**
令 y=x&(x-1)，则 y 为将 x 的最低设置位从 1 变成 0 之后的数，显然 0≤y<x0, bits[x]=bits[y]+1。因此对任意正整数 x，都有 bits[x]=bits[x&(x−1)]+1
 */
func countBits2(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// y = x & (x-1)，其中y表示将x的最低位从1变为0，
		// 那么 x 中 1 的个数一定等于 y 中 1 的个数加1，
		// 假设 bits[x] 表示x中含有1的个数，那么 bits[x] = bits[y] + 1, 也就是：bits[x] = bits[x&(x-1)] + 1
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}
