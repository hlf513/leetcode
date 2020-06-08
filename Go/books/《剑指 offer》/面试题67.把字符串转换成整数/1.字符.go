package main

import "strings"

/**
解题思路：注意边界：超出范围，负数，空格，有字符串
时间复杂度：N
空间复杂度：N
*/
func strToInt(str string) int {
	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return 0
	}

	// 保留符号的 str
	s0 := str
	// 过滤符号
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
	}
	// 声明返回值
	n := 0
	for _, b := range []byte(str) { // 按字节处理
		// 用于判断是否是数字
		b -= '0'
		// 不是数字
		if b > 9 {
			break
		}
		n = n*10 + int(b) //  十进制，进位
		// 判断越界
		if n > 2147483647 {
			// 表示已溢出
			n = 2147483648
			break
		}
	}
	// 判断符号 && 越界
	if s0[0] == '-' {
		n = -n
	} else if n >= 2147483648 {
		// 修正正数范围
		n = 2147483647
	}

	return n
}
