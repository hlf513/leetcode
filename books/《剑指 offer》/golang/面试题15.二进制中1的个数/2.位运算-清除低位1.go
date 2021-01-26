package main

/**
题目：请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数
解题思路： 位运算，清除低位1
时间复杂度：M (1的个数)
空间复杂度：1
*/
func hammingWeight(num uint32) int {
	var ret int
	for num != 0 {
		ret++
		num = num & (num - 1)
	}
	return ret
}
