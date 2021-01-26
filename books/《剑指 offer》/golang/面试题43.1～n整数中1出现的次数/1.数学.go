package main

import "fmt"

/**
题目：给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
解题思路：找规则，n 分为三部分 high、cur、low；digit 表示 cur 所在位，每次向左移位一个，根据公式计算结果
	- cur = 0, res += high * digit
	- cur = 1, res += high * digit + low + 1
	- cur > 1, res += high * digit + digit
时间复杂度：logn
空间复杂度：1
*/
func countDigitOne(n int) int {
	var res int
	// 定义公式中的变量
	digit := 1
	high := n / 10
	cur := n % 10
	low := 0
	// 开始遍历每位，计算 result
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += high*digit + digit
		}
		// 更新公式中的变量
		low += cur * digit
		cur = high % 10
		high = high / 10
		digit *= 10
	}

	return res
}

func main() {
	fmt.Println(countDigitOne(12))
}
