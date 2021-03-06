package main

/**
题目：一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。答案需要取模 1e9+7（1000000007）
解题思路：类似斐波数列，注意 n=0时，a=1，而不是0
时间复杂度：N
空间复杂度：N
*/
func numWays(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}
