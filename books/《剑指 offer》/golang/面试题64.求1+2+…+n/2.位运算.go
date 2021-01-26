package main

/**
题目：不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
解题思路：位运算，关键是 a*b 如何计算
时间复杂度：快速乘 O(logn)
空间复杂度：1
*/
func sumNums2(n int) int {
	return quickMulti(n, n+1) >> 1
}

// quickMulti 计算 a*b (俄罗斯农民乘法)
func quickMulti(a, b int) int {
	var ans = 0
	for ; b > 0; b >>= 1 {
		if b&1 == 1 {
			ans += a
		}
		a <<= 1
	}

	return ans
}
