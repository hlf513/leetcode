package main

import "fmt"

func main() {
	s := "  hello world!  "
	s = "a"
	fmt.Println(reverseWords(s))
}

/**
题目： 反转字符串中的单词
解题思路：使用双指针分隔单词
时间复杂度：N
空间复杂度：N
*/
func reverseWords(s string) string {
	// 切割单词
	var i, j = len(s) - 1, len(s) - 1
	var tmp []string
	for i >= 0 { // =0 因为边界条件（一个字符）
		j = i
		// 找到第一个非空字符
		for i >= 0 && s[i] == ' ' {
			i--
			j-- // j 在单词的最后边
		}
		// i 移动到单词的最左边
		for i >= 0 && s[i] != ' ' {
			i--
		}
		// 保存单词
		if s[i+1] != ' ' { // 过滤单词前的空格
			tmp = append(tmp, s[i+1:j+1])
		}
	}
	if len(tmp) == 0 {
		return ""
	}
	// 反转单词
	var result string
	for i := 0; i < len(tmp); i++ {
		result += " " + tmp[i]
	}
	return result[1:]
}
