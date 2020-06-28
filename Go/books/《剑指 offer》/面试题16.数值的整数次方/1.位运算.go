package main

/**
题目：实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。
解题思路：位运算，把 n 二进制化，0 则 *x，1则 *x 后放入 res，注意处理负数
时间复杂度：
空间复杂度：
*/
func myPow(x float64, n int) float64 {
	// 边界处理
	if x == 0 {
		return 0
	}
	res := 1.0
	// 处理负数
	if n < 0 {
		x, n = 1/x, -n
	}
	// 开始计算：把 n 二进制化，0 则 *x，1则 *x 后放入 res
	for n != 0 {
		if n&1 == 1 { // 放入 res
			res *= x
		}
		x *= x
		n >>= 1 // 右移
	}
	return res
}
