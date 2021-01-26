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
 *
 * 关键点：prefix
 * - 开始 prefix = strs[0]，若满足条件则返回，否则删除最后一个字符，继续匹配
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
