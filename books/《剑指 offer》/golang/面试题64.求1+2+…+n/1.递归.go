package main

/**
题目：不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
解题思路：递归，注意退出条件
时间复杂度：N
空间复杂度：N
*/
func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	// 定义求和函数
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1) // 退出条件: n <=0
	}
	sum(n)
	return ans
}
