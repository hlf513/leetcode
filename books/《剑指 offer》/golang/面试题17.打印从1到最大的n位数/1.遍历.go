package main

/**
题目：输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
解题思路：求 10 的 n 次方
时间复杂度：10^n
空间复杂度：1
*/
func printNumbers(n int) []int {
	// 求10的 n 次方
	max := 1
	n++
	for n > 1 {
		max *= 10
		n--
	}

	// 遍历输出
	var result []int
	for i := 1; i < max; i++ {
		result = append(result, i)
	}

	return result
}
