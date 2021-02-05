package main

import (
	"fmt"
	"math"
)

func main() {

	var s = "abcabcbb"

	//fmt.Println(bruteForce(s))
	fmt.Println(slidingWindow(s))
	//fmt.Println(slidingWindowOptimized(s))
}

/**
 * 暴力法
 * 解题思路： 获取所有子串，检查子串是否无重复字符
 * 时间复杂度：O(n^3)
 * 空间复杂度：O(?)
*/
func bruteForce(s string) int {

	return 0
}

/**
 * 滑动窗口
 * 解题思路：两个指针 + 集合；
 *  	1. 若右值在集合里：右指针+1（右值写入集合）；计算最大长度
 * 		2. 若右值不在集合里：左指针+1（元素删除）
 * 时间复杂度：O(2n)
 * 空间复杂度：O(?)
*/
func slidingWindow(s string) int {
	var n = len(s)
	var m = make(map[string]int)
	var i, j int
	var max float64

	for i < n && j < n {
		if _, ok := m[string(s[j])]; !ok {
			m[string(s[j])] = 1
			j++
			max = math.Max(max, float64(j-i))
		} else {
			delete(m, string(s[i]))
			i++
		}
	}

	return int(max)
}

/**
 * 解题思路：滑动窗口优化，若右值不在集合中，左指针+j
 * 时间复杂度：O(n)
 * 空间复杂度：O(?)
*/
func slidingWindowOptimized(s string) int {

	return 0
}
