package main

import (
	"fmt"
)

func main() {
	n, k := 4, 9

	fmt.Println(getPermutation(n, k))
}

/**
 * 注意 k = k-1 ；因为下标从0开始
 *
 * index =  k / (n-1)!
 * k' = k % (n-1)!
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
