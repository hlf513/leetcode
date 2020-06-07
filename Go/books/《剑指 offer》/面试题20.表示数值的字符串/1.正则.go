package main

import (
	"regexp"
	"strings"
)

func main() {

}

/**
题目：判断字符串是否表示数值（包括整数和小数）
解题思路：分析可以表达数值的字符串：+100"、"5e2"、"-123"、"3.1416"、"0123"
时间复杂度：N
空间复杂度：1
*/
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	res, _ := regexp.MatchString("^(([\\+\\-]?[0-9]+(\\.[0-9]*)?)|([\\+\\-]?\\.?[0-9]+))(e[\\+\\-]?[0-9]+)?$", s)

	return res
}
