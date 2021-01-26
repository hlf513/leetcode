package main

import "fmt"

/**
题目：0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。
解题思路：先推导出公式 f(n,m)=(m+x)%n，再递归求解
时间复杂度：
空间复杂度：
*/
func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	x := lastRemaining(n-1, m)
	return (m + x) % n // 删除后剩余的最后一个数字
}

func main() {
	fmt.Println(lastRemaining(5, 3))
}
