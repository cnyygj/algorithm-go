package dp

import "fmt"

/**
给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。

输入: num = 1775(110111011112)
输出: 8

输入: num = 7(01112)
输出: 4

 */


/**
方法一：
currentMaxLen[i] 表示包以第i位结尾的从num二进制低位至第i位连续1的最长长度
reverseMaxLen[i] 表示包含以第i位结尾的从低位到第i位最多翻转1个0->1 的连续1的最长长度

用num[i]表示整数num第i位的值

当num[i]=1时，currentMaxLen[i] = current[i-1]+1,currentMaxLen[i-1]一定包含i-1位，也就是和第i位连续，所以前i-1的最大长度连上第i位的长度就等于currentMaxLen[i],同理reverseMaxLen[i] = reverseMaxLen[i-1]+1;
num[i]=0时，连续中断，currentMaxLen[i]=0,而reverseMaxLen[i]允许翻转1次，但是reverseMaxLen[i]又必须包含第i位，也就是说只能翻转第i位，所以前面不能出现翻转，必须全是1，这个长度恰好就是currentMaxLen[i-1]，所以reverseMaxLen[i] = currentMaxLen[i-1]+1
遍历num所有位数，也就是32位后，reverseMaxLen数组中的最大值就是答案。

状态方程：
currentMaxLen[i] = num[i]==1?currentMaxLen[i-1]+1:0
reverseMaxLen[i] = num[i]==1?reverseMaxLen[i-1]+1:currentMaxLen[i-1]+1

 */
func reverseBits(num int) int {

	if num == 0 {
		return 1
	}

	var maxLen int
	// currentMaxLen[i]: 表示包以第i位结尾的从num二进制低位至第i位连续1的最长长度
	currentMaxLen := make(map[int]int, 32)
	fmt.Println(currentMaxLen)
	// reverseMaxLen[i]: 表示包含以第i位结尾的从低位到第i位最多翻转1个0->1 的连续1的最长长度
	reverseMaxLen := make(map[int]int, 32)

	for i := 0; i < 32; i++ {
		fmt.Println(num)
		// 判断当前整数的最低位是否是1
		if (num & 1) == 1 {
			currentMaxLen[i] = currentMaxLen[i-1] + 1
			reverseMaxLen[i] = reverseMaxLen[i-1] + 1

		} else {
			// 如果当前最低位是0，则重新更新低位至当前位连续1的最长长度
			currentMaxLen[i] = 0
			// 可将当前位翻转为1，那么当前位1的最长度 = i-1位置1的最长度+1
			reverseMaxLen[i] = currentMaxLen[i-1] + 1
		}

		maxLen = max(reverseMaxLen[i], maxLen)
		if num == 0 {
			return maxLen
		}

		num >>= 1
	}

	return maxLen
}

/**
方法2：
观察方法一状态方程，我们发现current和reverse第i位只和第i-1位有关，所以可以把动态数组优化成两个变量current和reverse，同时更新最大值max并作为结果返回。

 */
func reverseBits2(num int) int {

	if num == 0 {
		return 1
	}

	var maxLen int
	currentMaxLen := 0
	reverseMaxLen := 0

	for i := 0; i < 32; i++ {
		fmt.Println(num)
		// 判断当前整数的最低位是否是1
		if (num & 1) == 1 {
			currentMaxLen++
			reverseMaxLen++
		} else {
			reverseMaxLen = currentMaxLen + 1
			currentMaxLen = 0
		}

		maxLen = max(reverseMaxLen, maxLen)
		if num == 0 {
			return maxLen
		}

		num >>= 1
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
