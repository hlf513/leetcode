package main

import (
	"fmt"
)

func main() {
	n, k := 4, 9

	fmt.Println(getPermutation(n, k))
}

/**
 * 解题思路： 按位找规律
 * 		核心思想：把字符串分组，k/(n-1)! 表示在第几组的第几位
 * 		核心公式：
 * 		- index = k / (n-1)! 求字符串的下标
 * 		- k' = k % (n-1)! 更新最新字符串的 k 值
 * 		注意：因为下标从0开始，k 需要初始化为 k-1
 * 时间复杂度：O(n)
 * 空间复杂度：O(n)
 */
func getPermutation(n int, k int) string {
	// 求阶乘
	j := func(m int) int {
		var res = m
		for m > 1 {
			res = res * (m - 1)
			m--
		}
		return res
	}

	var nums = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	// 移除 slice 中的元素
	remove := func(i int) {
		nums = append(nums[:i], nums[i+1:]...)
	}

	var res string
	k = k - 1
	// 按照长度迭代
	for i := 0; i < n; i++ {
		var index int
		if i == n-1 {
			index = 0
		} else {
			index = k / j(n-i-1)
			k = k % j(n-i-1)
		}
		res += nums[index]
		//fmt.Println(nums, index, k)
		remove(index)
	}

	return res
}
