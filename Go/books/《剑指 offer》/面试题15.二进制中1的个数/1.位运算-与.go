package main

/**
题目：请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数
解题思路：位运算，利用 x & 1 = 1 和 x >> 1 累积1的位数
时间复杂度：logN(每次移位一个)
空间复杂度：1
*/
func hammingWeight2(num uint32) int {
	var ret uint32
	for num != 0 {
		ret += num & 1 // +0 无所谓
		num >>= 1      // num 右移一位
	}

	return int(ret)
}
