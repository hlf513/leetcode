package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "the sky  is blue"
	//var s =" "

	fmt.Println(reverseWords(s))
}

/**
 解题思路： 切分翻转
	1. 切分字符串
	2. 翻转字符串
 时间复杂度：O(n)
 空间复杂度：O(字符串长度)
*/
func reverseWords(s string) string {
	tmp := strings.Split(strings.Trim(s, " "), " ")
	n := len(tmp)

	var r []string
	for i := n - 1; i >= 0; i-- {
		if tmp[i] != "" {
			r = append(r, tmp[i])
		}
	}

	return strings.Join(r, " ")
}
