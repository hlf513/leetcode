package main

import "fmt"

func main() {
	ret := replaceSpace("ab x")
	fmt.Println(ret)
}

/**
题目：替换空格为%20
解题思路：1字符=>3字符、构建字符数组
时间复杂度：N
空间复杂度：N (s 的3倍）
*/
func replaceSpace(s string) string {
	var result []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result = append(result, '%', '2', '0')
		} else {
			result = append(result, s[i])
		}
	}

	return string(result)
}
