package main

import "fmt"

func main() {
	fmt.Println(translateNum(12258))
}

/**
题目：
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
解题思路：抽象为二叉树，以最后一位为根节点，每次只有两个分支可选，问题转化为从根结点到叶子节点的路径个数
时间复杂度：NlogN
空间复杂度：N
*/
func translateNum(num int) int {
	s := fmt.Sprintf("%d", num)
	return backtrace(s, 0)
}

func backtrace(s string, p int) int {
	// 退出条件
	n := len(s)
	if p == n {
		return 1
	}

	// num[i]和num[i−1]不能合成一个字符
	if p == n-1 || s[p] == '0' || s[p:p+2] > "25" {
		return backtrace(s, p+1)
	}

	// num[i]和num[i−1]能合成一个字符
	return backtrace(s, p+1) + backtrace(s, p+2)
}
