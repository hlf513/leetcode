package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "anagram"
	t := "nbgbrbm"

	fmt.Println(isAnagram2(s, t))
}

/**
方法一：字符排序
方法二：统计字符出现次数
*/
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sa := strings.Split(s, "")
	ta := strings.Split(t, "")
	sort.Strings(sa)
	sort.Strings(ta)
	saa := strings.Join(sa, "")
	sbb := strings.Join(ta, "")

	return saa == sbb
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sm = map[string]int{}
	var tm = map[string]int{}
	for i := 0; i < len(s); i++ {
		sm[string(s[i])]++
	}
	for i := 0; i < len(t); i++ {
		tm[string(t[i])]++
	}
	if len(sm) != len(tm) {
		return false
	}

	//fmt.Println(tm, sm)
	for k, v := range sm {
		if v1, _ := tm[k]; v1 != v {
			return false
		}
	}

	return true
}
