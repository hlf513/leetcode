package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
题目：数字以0123456789101112131415…的格式序列化到一个字符序列中。
解题思路：
	1. 先找 n 对应的数字位数
	2. 再找 n 对应的数字
	3. 最后找 n 对应的数字的位置
时间复杂度：
空间复杂度：
*/
func findNthDigit(n int) int {
	// digit，n 的位数
	// start，digit 的起始数字（1，10，100...）
	// count，n 的总字符数量
	var digit, start, count = 1, 1, 9
	// 1. 先找 n 对应的数字位数
	for n > count {
		n -= count
		start *= 10
		digit++
		count = 9 * start * digit
	}
	// 2. 再找 n 对应的数字
	num := start + (n-1)/digit
	// 3. 最后找 n 对应的数字的位置
	s := strings.Split(fmt.Sprintf("%d", num), "")
	index := (n - 1) % digit

	ret, _ := strconv.Atoi(s[index])

	return ret
}

func main() {
	fmt.Println(findNthDigit(3))
}
