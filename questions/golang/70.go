package main

import "fmt"

func main() {
	fmt.Println(climbStairs2(44))
}

/**
方法一：斐波那契数列（递归）
方法二：斐波那契数列（迭代）
*/
// 方法一（leetcode 会超时）
func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

// 方法二
func climbStairs2(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	m := make(map[int]int)
	m[0] = 1
	m[1] = 2
	for i := 2; i < n; i++ {
		m[i] = m[i-1] + m[i-2]
	}
	return m[n-1]
}
