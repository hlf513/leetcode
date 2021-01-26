package main

/**
题目：写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项；答案需要取模 1e9+7（1000000007）
解题思路：根据公式作答
时间复杂度：N
空间复杂度：N
*/
func fib(n int) int {
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	for i := 2; i <= n; i++ {
		cache[i] = (cache[i-1] + cache[i-2]) % 1000000007
	}

	return cache[n]
}
