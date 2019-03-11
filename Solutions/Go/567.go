package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s1 = "ab"
	var s2 = "eidboaooo"

	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {

	count1 := len(s1)
	count2 := len(s2)
	if count2 < count1 {
		return false
	}
	// 获取所有的子串
	for i := 0; i <= count2-count1; i++ {
		// 判断是否符合条件
		//fmt.Println(i, s2[i:i+count1])
		if reflect.DeepEqual(calCount(s2[i:i+count1]), calCount(s1)) {
			return true
		}
	}

	return false
}

/**
 * 计算字符串中字符出现的次数
 */
func calCount(s string) map[string]int {
	var m = make(map[string]int)

	n := len(s)
	for i := 0; i < n; i++ {
		m[string(s[i])]++
	}

	return m
}
