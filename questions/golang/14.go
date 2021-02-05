package main

import (
	"fmt"
	"strings"
)

func main() {

	//var strs = []string{"flower", "flow", "flight"}
	//var strs = []string{"a"}
	var strs = []string{""}

	fmt.Println(longestcommonPrefix(strs))

}

/**
 * 水平扫描（右左）
 * 解题思路：构建 prefix，循环匹配元素，若不匹配，从右往左依次递减
 * 		关键点：prefix
 * 		- 开始 prefix = strs[0]，若满足条件则返回，否则删除最后一个字符，继续匹配
 * 问题：若最后一个元素字符串很短，则有很多无效的比较
 */
func longestcommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix = strs[0]
	for _, v := range strs {
		for strings.Index(v, prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}


/**
其他解法
1. 水平扫描（左=>右）；构建 prefix，循环匹配元素，若不匹配，从左往右依次递增
1. 分治（元素）；按元素个数分治
1. 分治（字符串）；按每个元素的字符串长度分治，每次抛弃右半区
*/