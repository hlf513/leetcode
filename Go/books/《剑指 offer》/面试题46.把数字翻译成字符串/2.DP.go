package main

import "fmt"

/**
解题思路：和「爬楼梯」思路一样；
时间复杂度：NlogN
空间复杂度：N
*/
func translateNum2(num int) int {
	s := fmt.Sprintf("%d", num)
	var result = make(map[int]int)
	result[0] = 1
	result[1] = 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == '0' || s[i-1:i+1] > "25" { // num[i]和num[i−1]不能合成一个字符
			result[i+1] = result[i]
		} else { // num[i]和num[i−1]能合成一个字符
			result[i+1] = result[i] + result[i-1]
		}
	}

	return result[len(s)]
}
