package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s = "25525511135"

	fmt.Println(restoreIpAddresses(s))

}

func restoreIpAddresses(s string) []string {
	var res []string

	restore(s, 4, "", &res)

	return res
}

/**
递归
1. 把字符串分为5段（最后一段应该为空）
2. 判断每段是否符合条件，且最后一段是空
3. 满足条件则放入结果集

@param s 截断后的字符串
@param k 还剩下几段（一共5段）
@param out 满足条件后，拼接的字符串
@param res 保存所有的结果

时间复杂度：O(n^5)
空间复杂度：O(n)
*/
func restore(s string, k int, out string, res *[]string) {
	// 退出条件
	if k == 0 {
		if len(s) == 0 { // 若有剩余字符，则说明不符合条件
			out = strings.TrimRight(out, ".")
			*res = append(*res, out)
		}
		return
	}
	// 一共四段
	for i := 1; i < 4; i++ {
		if len(s) >= i && isValid(string(s[0:i])) { // len(s) >= i 防止越界
			// 符合条件，判断下一个
			restore(string(s[i:]), k-1, out+string(s[0:i])+".", res)
		}
	}
}

/**
判断给定的 s 是否满足 IP 段要求
*/
func isValid(s string) bool {
	if len(s) > 3 || s == "" || (len(s) > 1 && string(s[0]) == "0") {
		return false
	}
	i, _ := strconv.Atoi(s)
	return i <= 255 && i >= 0
}
